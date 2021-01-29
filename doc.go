//
//I want to be able to have a "declarative statement" of the harness configuration and topology so that I can generate terraform content from templates to account for the differences between clouds that terraform supports.
//
//Notes:
//
//	* given a yaml-based input, generate terraform files
//	* does not need to generate "all" terraform files needed for a harness, but could be used to generate some of them
//	* input yaml file format needs to describe:
//		* "configuration" of the harness - number of box's and dut's, networks, etc.
//		* type of network, number of ips allocate, etc.
//	* "topology" - how each box or dut is connected to each network (or not)
//

package gocookbook
