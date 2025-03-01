// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ips

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Remove() *cobra.Command {
	var reservationID string
	// removeIPCmd represents the removeIp command
	var removeIPCmd = &cobra.Command{
		Use:   "remove",
		Short: "Command to remove IP reservation.",
		Long: `Example:	

metal ip remove --id [reservation-UUID]

`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := c.ProjectService.Remove(reservationID)
			if err != nil {
				return errors.Wrap(err, "Could not remove IP address Reservation")
			}

			fmt.Println("IP reservation removed successfully.")
			return nil
		},
	}

	removeIPCmd.Flags().StringVarP(&reservationID, "id", "i", "", "UUID of the reservation")

	_ = removeIPCmd.MarkFlagRequired("id")
	return removeIPCmd
}
