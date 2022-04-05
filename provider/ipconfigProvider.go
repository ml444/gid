package provider

// 发布前配置所有服务节点IP的映射， 每个服务节点必须具有相同的映射
// 运行时每个服务节点根据本机IP取得IP映射中的值作为自己的机器号。

import (
	"github.com/ml444/gid/utils"
	"strings"
)

type IpConfigurableMachineIdProvider struct {
	machineId int
	ipsMap    map[string]int
}

func (p IpConfigurableMachineIdProvider) init(ips string) {
	ip := utils.GetHostIp()

	// if (ip == "") {
	// 	msg := "Fail to get host IP address. Stop to initialize the IpConfigurableMachineIdProvider provider.";

	// 	log.error(msg);
	// 	throw new IllegalStateException(msg)
	// }

	// if (!ipsMap.containsKey(ip)) {
	// 	String msg = String
	// 			.format("Fail to configure ID for host IP address %s. Stop to initialize the IpConfigurableMachineIdProvider provider.",
	// 					ip);

	// 	log.error(msg);
	// 	throw new IllegalStateException(msg)
	// }

	p.machineId = p.ipsMap[ip]

	// log.info("IpConfigurableMachineIdProvider.init ip {} id {}", ip,
	// 		machineId);

	p.setIp(ips)
}

func (p IpConfigurableMachineIdProvider) setIp(ips string) {
	// log.debug("IpConfigurableMachineIdProvider ips {}", ips);
	if len(ips) != 0 {
		ipArray := strings.Split(ips, ",")

		for i := 0; i < len(ipArray); i++ {
			p.ipsMap[ipArray[i]] = i
		}
	}
}

func (p IpConfigurableMachineIdProvider) GetMachineId() int {
	return p.machineId
}

func (p IpConfigurableMachineIdProvider) SetMachineId(machineId int) {
	p.machineId = machineId
}
