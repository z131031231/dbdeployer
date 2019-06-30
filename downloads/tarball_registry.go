// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2019 Giuseppe Maxia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package downloads

// This file was generated during build. Do not edit.
// Build time: Sun Jun  9 18:11:10 CEST 2019

var DefaultTarballRegistry = TarballCollection{
	// Version of dbdeployer when the list was last updated
	DbdeployerVersion: "1.32.0",
	Tarballs: []TarballDescription{

		{
			Name:            "tidb-master-darwin-amd64.tar.gz",
			Checksum:        "",
			OperatingSystem: "Darwin",
			Url:             "https://download.pingcap.org/tidb-master-darwin-amd64.tar.gz",
			Flavor:          "tidb",
			Minimal:         false,
			Size:            26485675,
			ShortVersion:    "3.0",
			Version:         "3.0.0",
			Notes:           "No checksum until the file name includes the version",
		},
		{
			Name:            "tidb-master-linux-amd64.tar.gz",
			Checksum:        "",
			OperatingSystem: "Linux",
			Url:             "https://download.pingcap.org/tidb-master-linux-amd64.tar.gz",
			Flavor:          "tidb",
			Minimal:         false,
			Size:            25918765,
			ShortVersion:    "3.0",
			Version:         "3.0.0",
			Notes:           "No checksum until the file name includes the version",
		},
		{
			Name:            "mysql-5.7.26-macos10.14-x86_64.tar.gz",
			Checksum:        "SHA512:ae84b0cfe3cf274fc79adb3db03b764d47033aea970cc26edcdd4adbe5b2e3d28bf4f98f2ee321f16e788d69cbe3a08bf39fa5329d8d7a67bee928d964891ed8",
			OperatingSystem: "Darwin",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-5.7.26-macos10.14-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            337246370,
			ShortVersion:    "5.7",
			Version:         "5.7.26",

			UpdatedBy: "dbdeployer",
		},
		{
			Name:            "mysql-8.0.16-macos10.14-x86_64.tar.gz",
			Checksum:        "SHA512:30fb86c929ad1f384622277dbc3d686f5530953a8f7e2c7adeb183768db69464e93a46b4a0ec212d006e069f1b93db0bd0a51918eaa7e3697ea227d86082d892",
			OperatingSystem: "Darwin",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.16-macos10.14-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            153205937,
			ShortVersion:    "8.0",
			Version:         "8.0.16",
		},
		{
			Name:            "mysql-8.0.15-macos10.14-x86_64.tar.gz",
			Checksum:        "SHA512:14ea13a9529bded717040699be3cc6ad7a88193ea9453acc36ef84fce7e74b6ef9d8e6953c61836a92b0f81eb96c54e13f18b1cae01dcd09520073c15734499c",
			OperatingSystem: "Darwin",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-8.0.15-macos10.14-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            138725124,
			ShortVersion:    "8.0",
			Version:         "8.0.15",
		},
		{
			Name:            "mysql-5.7.25-macos10.14-x86_64.tar.gz",
			Checksum:        "SHA512:8ae57e1fc057871f993085cc7e061d2b9f56df80c49f4b813e180fe446f318c4fcb80c41256597335d43e3bc39979eebdb113218eaaa2a1e45c4c5694a8bd883",
			OperatingSystem: "Darwin",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.7.25-macos10.14-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            337018862,
			ShortVersion:    "5.7",
			Version:         "5.7.25",
		},
		{
			Name:            "mysql-5.6.41-macos10.13-x86_64.tar.gz",
			Checksum:        "SHA512:9bce235b962d51afb25490a0e3d8fe4eab5f973083f467c7149779357b94ea82e15dca6f263285a6ebd6d52f38febc6e7dca042f285809e6eea5dde6e4638a9d",
			OperatingSystem: "Darwin",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.6.41-macos10.13-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            175713012,
			ShortVersion:    "5.6",
			Version:         "5.6.41",
		},
		{
			Name:            "mysql-5.5.53-osx10.9-x86_64.tar.gz",
			Checksum:        "SHA512:ce466d6ccdb8250783e763955dd2400826345d7c9df4d317981fa947256eedb8ecb4376d90847351e8994d43775bd56e8a14f09e6a448a352508b2de1b7a755c",
			OperatingSystem: "Darwin",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.5.53-osx10.9-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            114490251,
			ShortVersion:    "5.5",
			Version:         "5.5.53",
		},
		{
			Name:            "mysql-5.1.73-osx10.6-x86_64.tar.gz",
			Checksum:        "SHA512:9e21642d009ac736607d38b3b4dd0b566db39995dbcf1515cd73026c5362419dcf7d085407c3f89aab59bbf740e383cf6a257517ec00653ecd7382f3166bca38",
			OperatingSystem: "Darwin",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.1.73-osx10.6-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            81598665,
			ShortVersion:    "5.1",
			Version:         "5.1.73",
		},
		{
			Name:            "mysql-5.0.96-osx10.5-x86_64.tar.gz",
			Checksum:        "SHA512:a5f3959732fe8b4a9cb3b41d7c2f7d869cae38ffc844989b7c98cdbf7496e6d65ea1794819d2a6aea8dd13340252a8c553dcdfbc904415ecd13adb493aa5ae68",
			OperatingSystem: "Darwin",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.0.96-osx10.5-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            60958921,
			ShortVersion:    "5.0",
			Version:         "5.0.96",
		},
		{
			Name:            "mysql-8.0.16-linux-glibc2.12-x86_64.tar.xz",
			Checksum:        "MD5: 60d18d1b324104c83da33dcd7a989816",
			OperatingSystem: "Linux",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.16-linux-glibc2.12-x86_64.tar.xz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            460733332,
			ShortVersion:    "8.0",
			Version:         "8.0.16",
		},
		{
			Name:            "mysql-8.0.16-linux-x86_64-minimal.tar.xz",
			Checksum:        "MD5: 7bac88f47e648bf9a38e7886e12d1ec5",
			OperatingSystem: "Linux",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.16-linux-x86_64-minimal.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            44246564,
			ShortVersion:    "8.0",
			Version:         "8.0.16",
		},
		{
			Name:            "mysql-5.7.26-linux-glibc2.12-x86_64.tar.gz",
			Checksum:        "SHA512:abc2e35bbe28eb1651ae0f4cbd2c2febe1c3e0b6f510239f64df40d4a50f860e419e2570440b35255e8a471050358abd7fef87606f74b55eb2700678b3949244",
			OperatingSystem: "Linux",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-5.7.26-linux-glibc2.12-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            644869837,
			ShortVersion:    "5.7",
			Version:         "5.7.26",
		},
		{
			Name:            "mysql-5.6.44-linux-glibc2.12-x86_64.tar.gz",
			Checksum:        "SHA512:cd638b1f087cee7cda0f80361917b377b736b83d8aaf7c15db39b9aae3ea143ffb1cd68c9e30fea8e774f6fb48de2df0878e56a4e9be4559142ed3dfae36dbb0",
			OperatingSystem: "Linux",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-5.6/mysql-5.6.44-linux-glibc2.12-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            329105487,
			ShortVersion:    "5.6",
			Version:         "5.6.44",
		},
		{
			Name:            "mysql-5.5.62-linux-glibc2.12-x86_64.tar.gz",
			Checksum:        "SHA512:7bf9d740a5e702ba53421b5b01a9d26e86e7cdc7fe8fa86871683c61be6395227bac6547136a9b080f4f918d249324e6754639924d04b89877a81496e6a52e4b",
			OperatingSystem: "Linux",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-5.5/mysql-5.5.62-linux-glibc2.12-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            198993245,
			ShortVersion:    "5.5",
			Version:         "5.5.62",
		},
		{
			Name:            "mysql-8.0.15-linux-glibc2.12-x86_64.tar.xz",
			Checksum:        "SHA512:7bcb6e1b43f67db9d1bb323ef923ee9c12be072b03e8e09e06a68802ffe320c1eb8e3f2bf4f4c15f549b0e9048809fc366eab103890d1de7284451cd74060571",
			OperatingSystem: "Linux",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-8.0.15-linux-glibc2.12-x86_64.tar.xz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            376303592,
			ShortVersion:    "8.0",
			Version:         "8.0.15",
		},
		{
			Name:            "mysql-8.0.13-linux-glibc2.12-x86_64.tar.xz",
			Checksum:        "SHA512:a15cbae5ba6e7f908da528143549d607cedd3970e579bf22b561942105349a5b20b289950a5ae3c2ef569b89e2a46ed58da50981fa8e6604b64e4224e5b283d3",
			OperatingSystem: "Linux",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-8.0.13-linux-glibc2.12-x86_64.tar.xz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            393852364,
			ShortVersion:    "8.0",
			Version:         "8.0.13",
		},
		{
			Name:            "mysql-5.7.25-linux-glibc2.12-x86_64.tar.gz",
			Checksum:        "SHA512:7ec7036ac83e4995668ce92fc083ce41e1b090da66074b6734e413a6ba9d2dc1ad66869a06f2a8a461bc238b960d36e3afb8b3cc5b1a3088796cd9b246af7a7f",
			OperatingSystem: "Linux",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.7.25-linux-glibc2.12-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            644862820,
			ShortVersion:    "5.7",
			Version:         "5.7.25",
		},
		{
			Name:            "mysql-5.6.43-linux-glibc2.12-x86_64.tar.gz",
			Checksum:        "SHA512:a4cab03d312940a6cc080c486973836fac7982837c2305e389d97983d6608cdeed5fd48bd1d6543f14fb8e77d553243dd71899f5a74fd736b77f46a51eeb42a2",
			OperatingSystem: "Linux",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.6.43-linux-glibc2.12-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            328740156,
			ShortVersion:    "5.6",
			Version:         "5.6.43",
		},
		{
			Name:            "mysql-5.5.61-linux-glibc2.12-x86_64.tar.gz",
			Checksum:        "SHA512:83723300168195f47bb3791cd3baefe8172570b9d3317525efb5970e19c54b72de80430fd79a3e1e6a81795daa9a1a9bdc3ed61e7b0b07b10915825b8767fb30",
			OperatingSystem: "Linux",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.5.61-linux-glibc2.12-x86_64.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            198630181,
			ShortVersion:    "5.5",
			Version:         "5.5.61",
		},
		{
			Name:            "mysql-5.1.73-linux-x86_64-glibc23.tar.gz",
			Checksum:        "SHA512:045cc42a1078aae77d2ba9d9e35d69df79feb33de3dc112095e669d5b298ac24ba963ac071ab7a0d3cbe917c68624c184e4d7c69c76dbf43c4df32cb2e4da1b2",
			OperatingSystem: "Linux",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.1.73-linux-x86_64-glibc23.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            133630298,
			ShortVersion:    "5.1",
			Version:         "5.1.73",
		},
		{
			Name:            "mysql-5.0.96.tar.xz",
			Checksum:        "SHA512:c39621e43778cc971dc979a065b283864edcb15eb3aced59b5a596dbb56936a6f78feb141d5fb1e0213f130092cff1d93613c0b2beee8b8531796a745ad2bf4b",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-5.0.96.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            5464684,
			ShortVersion:    "5.0",
			Version:         "5.0.96",
		},
		{
			Name:            "mysql-5.1.72.tar.xz",
			Checksum:        "SHA512:6fc8ad76a831fc22f9b7bffed8fb4eb4575a122cebca8f5bab52dbe6454372b2f5681eb125141c2b1b5c547496e8d96686c83bdf3c4fdc390bd3614cd1060b9c",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-5.1.72.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            10390324,
			ShortVersion:    "5.1",
			Version:         "5.1.72",
		},
		{
			Name:            "mysql-5.5.61.tar.xz",
			Checksum:        "SHA512:0cc441cc1f9ddc35cc8783aeeee3ac319d95f3aec184703a2c9197577e23899c4e102010bcc1e60809c0d780ccec8d78e387f6afecb46ad009d9a94b20e9ebe6",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-5.5.61.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            6575380,
			ShortVersion:    "5.5",
			Version:         "5.5.61",
		},
		{
			Name:            "mysql-5.5.62.tar.xz",
			Checksum:        "SHA512:35b5decf5bb8b7291edba521efade90f8caeead9b52f17755ced98870a5e2652b53d00d9bd5f0911c9b1b21093119850874882e41bac58af7bd16e100de94c6d",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-5.5.62.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            6622968,
			ShortVersion:    "5.5",
			Version:         "5.5.62",
		},
		{
			Name:            "mysql-5.6.43.tar.xz",
			Checksum:        "SHA512:43934b37d87bed81cdb619ada12d1007e0fb787ff43bfe622608b61762c4e06da93a78bfbc14f2606039762f75fd9a7cb29e96f36792273772a47fe50b6886a8",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-5.6.43.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            8965164,
			ShortVersion:    "5.6",
			Version:         "5.6.43",
		},
		{
			Name:            "mysql-5.6.44.tar.xz",
			Checksum:        "SHA512:23cb2f8d289796b253c87ec3c171d1e8b1533c89284fa7ecad6f52fb8e90c46024984d562e056380cdbc9664ce284973121a2a9ae328ca4a8c599f1fa8eaf30f",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-5.6.44.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            9126724,
			ShortVersion:    "5.6",
			Version:         "5.6.44",
		},
		{
			Name:            "mysql-5.7.25.tar.xz",
			Checksum:        "SHA512:984f5fbad834b09028cc0353a36bfc6b0d4b9398f68298f14472be56dd4d35b170fb3ed34a0c985558922b2e9586af4a0d629c59731fd23ae171953cc9fe1bff",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-5.7.25.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            23274552,
			ShortVersion:    "5.7",
			Version:         "5.7.25",
		},
		{
			Name:            "mysql-5.7.26.tar.xz",
			Checksum:        "SHA512:b1642e61ad60f157d3b67e3372319b8cc1b057893dc7a18ad0d332c516fd5f3b78b4f763485039715ae7bfa33f7e36410c8331ad1dcd22164b4fba57993790b8",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-5.7.26.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            23274552,
			ShortVersion:    "5.7",
			Version:         "5.7.26",
		},
		{
			Name:            "mysql-5.0.96-linux-x86_64-glibc23.tar.gz",
			Checksum:        "SHA512:b8ee74710ea132253bf435844120daa22c04c098005f7cb3137e0f95b91a50d303fcc2ab58fdaf59af2fcb13148342dac3ec941c49eb44f6d89dcde2e1bc7707",
			OperatingSystem: "Linux",
			Url:             "https://downloads.mysql.com/archives/get/file/mysql-5.0.96-linux-x86_64-glibc23.tar.gz",
			Flavor:          "mysql",
			Minimal:         false,
			Size:            127124541,
			ShortVersion:    "5.0",
			Version:         "5.0.96",
		},
		{
			Name:            "mysql-4.1.22.tar.xz",
			Checksum:        "SHA512:6bd592bcccb6622be24b729861ed62a8830f2ca1bdc494256e2cb8bd2c3107732dc3e1ebeb4f3df6fe616d142e6f0ba92369d36f335a175d90c6b06eb79d4f31",
			OperatingSystem: "Linux",
			Url:             "https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/mysql-4.1.22.tar.xz",
			Flavor:          "mysql",
			Minimal:         true,
			Size:            4592864,
			ShortVersion:    "4.1",
			Version:         "4.1.22",
		},
		{
			Name:            "mysql-cluster-8.0.16-dmr-macos10.14-x86_64.tar.gz",
			Checksum:        "SHA512:bc80bce92778b7f9df93c1c466c829e657d2afc449aa54b92309b4a30d14d43f81c63a2a8fef164658627a1b79e42bdba34dfc0601a38d68fb2bed68210a771e",
			OperatingSystem: "Darwin",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-Cluster-8.0/mysql-cluster-8.0.16-dmr-macos10.14-x86_64.tar.gz",
			Flavor:          "ndb",
			Minimal:         false,
			Size:            252431350,
			ShortVersion:    "8.0",
			Version:         "8.0.16",
		},
		{
			Name:            "mysql-cluster-gpl-7.6.10-macos10.14-x86_64.tar.gz",
			Checksum:        "SHA512:62a1ac68a3884c3157dd63a25167f28e2eac120a3cad398797396d54e85c74a389d1c6bfd37d5de5b3d2653e0d557682b263e76ca333c0f12583968e4512b19d",
			OperatingSystem: "Darwin",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-Cluster-7.6/mysql-cluster-gpl-7.6.10-macos10.14-x86_64.tar.gz",
			Flavor:          "ndb",
			Minimal:         false,
			Size:            482157481,
			ShortVersion:    "7.6",
			Version:         "7.6.10",
		},
		{
			Name:            "mysql-cluster-gpl-7.6.10-linux-glibc2.12-x86_64.tar.gz",
			Checksum:        "SHA512:effb2a4ac4d4a850b45d5389106f1398dbb73c9feaf0288e91f7963e707a5688dff9f999005ad1c652352d9f2928244f4ec2f0eee6a2d1a023e9a0ea998b6cc4",
			OperatingSystem: "Linux",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-Cluster-7.6/mysql-cluster-gpl-7.6.10-linux-glibc2.12-x86_64.tar.gz",
			Flavor:          "ndb",
			Minimal:         false,
			Size:            915502140,
			ShortVersion:    "7.6",
			Version:         "7.6.10",
		},
		{
			Name:            "mysql-cluster-8.0.16-dmr-linux-glibc2.12-x86_64.tar.gz",
			Checksum:        "SHA512:a587a774cc7a8f6cbe295272f0e67869c5077b8fb56917e0dc2fa0ea1c91548c44bd406fcf900cc0e498f31bb7188197a3392aa0d7df8a08fa5e43901683e98a",
			OperatingSystem: "Linux",
			Url:             "https://dev.mysql.com/get/Downloads/MySQL-Cluster-8.0/mysql-cluster-8.0.16-dmr-linux-glibc2.12-x86_64.tar.gz",
			Flavor:          "ndb",
			Minimal:         false,
			Size:            1100516061,
			ShortVersion:    "8.0",
			Version:         "8.0.16",
		},
	},
}
