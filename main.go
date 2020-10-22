/*
 * The MIT License
 *
 * Copyright (c) 2020 gelleson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 *
 */

package main

import (
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/spf13/cobra"
	"os"
)

var root = &cobra.Command{
	Use: "redis",
	Run: func(cmd *cobra.Command, args []string) {
		redis := miniredis.NewMiniRedis()

		addr, err := cmd.Flags().GetString("addrs")

		if err != nil {
			fmt.Println(err)

			os.Exit(2)
		}

		if err = redis.StartAddr(addr); err != nil {
			fmt.Println(err)

			os.Exit(2)
		}

		if err = redis.Start(); err != nil {
			fmt.Println(err)

			os.Exit(2)
		}

		for {
			continue
		}
	},
}

func main() {
	root.PersistentFlags().String("addrs", ":6379", ":6379")

	if err := root.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
