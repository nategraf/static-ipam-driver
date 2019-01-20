package driver

import (
	"net"
	"reflect"

	"github.com/docker/go-plugins-helpers/ipam"
	"github.com/docker/libnetwork/types"
	"github.com/sirupsen/logrus"
)

type Driver struct{}

// StaticAddressSpace is the one address space understood by this driver.
const StaticAddrSpace = "static"

// unwrap gives the pointed to value if the i is an non-nil pointer.
func unwrap(i interface{}) interface{} {
	if v := reflect.ValueOf(i); v.Kind() == reflect.Ptr && !v.IsNil() {
		return v.Elem()
	}
	return i
}

// logRequest logs request inputs and results.
func logRequest(fname string, req interface{}, res interface{}, err error) {
	req, res = unwrap(req), unwrap(res)
	if err == nil {
		if res == nil {
			logrus.Infof("%s(%v)", fname, req)
		} else {
			logrus.Infof("%s(%v): %v", fname, req, res)
		}
		return
	}
	switch err.(type) {
	case types.MaskableError:
		logrus.WithError(err).Infof("[MaskableError] %s(%v): %v", fname, req, err)
	case types.RetryError:
		logrus.WithError(err).Infof("[RetryError] %s(%v): %v", fname, req, err)
	case types.BadRequestError:
		logrus.WithError(err).Warnf("[BadRequestError] %s(%v): %v", fname, req, err)
	case types.NotFoundError:
		logrus.WithError(err).Warnf("[NotFoundError] %s(%v): %v", fname, req, err)
	case types.ForbiddenError:
		logrus.WithError(err).Warnf("[ForbiddenError] %s(%v): %v", fname, req, err)
	case types.NoServiceError:
		logrus.WithError(err).Warnf("[NoServiceError] %s(%v): %v", fname, req, err)
	case types.NotImplementedError:
		logrus.WithError(err).Warnf("[NotImplementedError] %s(%v): %v", fname, req, err)
	case types.TimeoutError:
		logrus.WithError(err).Errorf("[TimeoutError] %s(%v): %v", fname, req, err)
	case types.InternalError:
		logrus.WithError(err).Errorf("[InternalError] %s(%v): %v", fname, req, err)
	default:
		// Unclassified errors should be treated as bad.
		logrus.WithError(err).Errorf("[UNKNOWN] %s(%v): %v", fname, req, err)
	}
}

func (d *Driver) GetDefaultAddressSpaces() (res *ipam.AddressSpacesResponse, err error) {
	defer func() { logRequest("GetDefaultAddressSpaces", nil, res, err) }()

	return &ipam.AddressSpacesResponse{
		LocalDefaultAddressSpace:  StaticAddrSpace,
		GlobalDefaultAddressSpace: StaticAddrSpace,
	}, nil
}

func (d *Driver) RequestPool(req *ipam.RequestPoolRequest) (res *ipam.RequestPoolResponse, err error) {
	defer func() { logRequest("RequestPool", req, res, err) }()

	if req.Pool == "" {
		return nil, ErrUnsupportedAllocation{}
	}

	if req.AddressSpace != StaticAddrSpace {
		return nil, ErrInvalidAddrSpace(req.AddressSpace)
	}

	return &ipam.RequestPoolResponse{
		PoolID: req.Pool,
		Pool:   req.Pool,
		Data:   nil,
	}, nil
}

func (d *Driver) ReleasePool(req *ipam.ReleasePoolRequest) (err error) {
	defer func() { logRequest("ReleasePool", req, nil, err) }()
	return nil
}

func (d *Driver) RequestAddress(req *ipam.RequestAddressRequest) (res *ipam.RequestAddressResponse, err error) {
	defer func() { logRequest("RequestAddress", req, res, err) }()

	_, subnet, err := net.ParseCIDR(req.PoolID)
	if err != nil {
		return nil, ErrParsePool(req.PoolID)
	}

	if req.Address != "" {
		subnet.IP = net.ParseIP(req.Address)
		if subnet.IP == nil {
			return nil, ErrParseIP(req.Address)
		}
	} else {
		// Obtain the first host address in the pool. (e.g. 192.168.0.0/24 -> 192.168.0.1)
		subnet.IP[len(subnet.IP)-1] |= 0x1
	}

	return &ipam.RequestAddressResponse{
		Address: subnet.String(),
		Data:    nil,
	}, nil
}

func (d *Driver) ReleaseAddress(req *ipam.ReleaseAddressRequest) (err error) {
	defer func() { logRequest("ReleaseAddress", req, nil, err) }()
	return nil
}

func (d *Driver) GetCapabilities() (res *ipam.CapabilitiesResponse, err error) {
	defer func() { logRequest("GetCapabilities", nil, res, err) }()
	return &ipam.CapabilitiesResponse{RequiresMACAddress: false}, nil
}
