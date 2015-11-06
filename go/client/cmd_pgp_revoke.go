// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol"
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

type CmdPGPRevoke struct {
	id string
}

func (c *CmdPGPRevoke) ParseArgv(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		return fmt.Errorf("Revoke takes exactly one key.")
	}
	c.id = ctx.Args()[0]
	return nil
}

func (c *CmdPGPRevoke) Run() (err error) {
	cli, err := GetRevokeClient()
	if err != nil {
		return err
	}

	protocols := []rpc.Protocol{
		NewSecretUIProtocol(G),
	}
	if err = RegisterProtocols(protocols); err != nil {
		return err
	}

	return cli.RevokeKey(context.TODO(), keybase1.RevokeKeyArg{
		KeyID: c.id,
	})
}

func NewCmdPGPRevoke(cl *libcmdline.CommandLine) cli.Command {
	return cli.Command{
		Name:         "revoke",
		ArgumentHelp: "<key-id>",
		Usage:        "Revoke a PGP key",
		Flags:        []cli.Flag{},
		Action: func(c *cli.Context) {
			cl.ChooseCommand(&CmdPGPRevoke{}, "revoke", c)
		},
	}
}

func (c *CmdPGPRevoke) GetUsage() libkb.Usage {
	return libkb.Usage{
		Config:     true,
		GpgKeyring: true,
		KbKeyring:  true,
		API:        true,
	}
}