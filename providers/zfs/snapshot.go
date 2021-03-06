package zfs

import (
	"errors"
	"net/url"

	"github.com/cerana/cerana/acomm"
	"github.com/cerana/cerana/zfs"
)

// SnapshotArgs are the arguments for the Snapshot handler.
type SnapshotArgs struct {
	Name      string `json:"name"`
	SnapName  string `json:"snapname"`
	Recursive bool   `json:"recursive"`
}

// Snapshot creates a snapshot of a filesystem or volume.
func (z *ZFS) Snapshot(req *acomm.Request) (interface{}, *url.URL, error) {
	var args SnapshotArgs
	if err := req.UnmarshalArgs(&args); err != nil {
		return nil, nil, err
	}

	if args.Name == "" {
		return nil, nil, errors.New("missing arg: name")
	}
	if args.SnapName == "" {
		return nil, nil, errors.New("missing arg: snapname")
	}

	ds, err := zfs.GetDataset(args.Name)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, ds.Snapshot(args.SnapName, args.Recursive)
}
