package test

import (
	"open_im_sdk/pkg/constant"
	"sync"
)

var LogLevel uint32 = 6
var PlatformID = int32(1)
var LogName = ""

var ReliabilityUserA = 1234567
var ReliabilityUserB = 1234567
var (
	TESTIP = "43.155.69.205"
	//TESTIP              = "121.37.25.71"
	APIADDR = "http://43.155.69.205" + ":10002"
	//WSADDR              = "wss://" + TESTIP + ":10001"
	//TESTIP  = "open-im-test.rentsoft.cn"
	//APIADDR = "https://" + TESTIP + ":50002"

	WSADDR              = "wss://" + TESTIP + ":50001"
	REGISTERADDR        = APIADDR + "/auth/user_register"
	TOKENADDR           = APIADDR + "/auth/user_token"
	SECRET              = "tuoyun"
	SENDINTERVAL        = 20
	GETSELFUSERINFO     = APIADDR + "/user/get_self_user_info"
	CREATEGROUP         = APIADDR + constant.CreateGroupRouter
	ACCOUNTCHECK        = APIADDR + "/user/account_check"
	GETGROUPSINFOROUTER = APIADDR + constant.GetGroupsInfoRouter
)

var coreMgrLock sync.RWMutex
var allLoginMgr map[int]*CoreNode

var allLoginMgrtmp []*CoreNode

var userLock sync.RWMutex

var allUserID []string
var allToken []string

//var allWs []*interaction.Ws
var sendSuccessCount, sendFailedCount int
var sendSuccessLock sync.RWMutex
var sendFailedLock sync.RWMutex

var msgNumInOneClient = 0

//var Msgwg sync.WaitGroup
var sendMsgClient = 0

var MaxNumGoroutine = 100000
