package common

import (
	"flag"
	"os"
	"path/filepath"
	"time"
)

var (
	SCRIPT_BIN_DIR, _   = filepath.Abs(filepath.Dir(os.Args[0]))
	WORKING_HOME_DIR, _ = filepath.Abs(filepath.Join(SCRIPT_BIN_DIR, ".."))
	LOG_HOME_DIR        = filepath.Join(WORKING_HOME_DIR, "logs")
	INVENTORY_HOME_DIR  = filepath.Join(WORKING_HOME_DIR, "inventory")
	CONFIG_HOME_DIR     = filepath.Join(INVENTORY_HOME_DIR, "config")
	CERT_HOME_DIR       = filepath.Join(INVENTORY_HOME_DIR, "cert")
	CGI_HOME_DIR        = filepath.Join(INVENTORY_HOME_DIR, "cgi")

	// cluster stuff
	ClusterGroup          string
	MemberMode            string
	OPLogMaxSeqGapToPull  uint64
	OPLogPullMaxCountOnce uint64
	OPLogPullInterval     time.Duration
	OPLogPullTimeout      time.Duration

	Host                           string
	CertFile, KeyFile              string
	Stage                          string
	ConfigHome, LogHome            string
	CpuProfileFile, MemProfileFile string
)

func init() {
	clusterGroup := flag.String("group", "default", "specify cluster group")
	memberMode := flag.String("mode", "read", "specify member mode (read or write)")
	opLogMaxSeqGapToPull := flag.Uint64("oplog_max_seq_gap_to_pull", 5,
		"specify max gap of sequnce of operation logs deciding whether to wait for missing operations or not")
	opLogPullMaxCountOnce := flag.Uint64("oplog_pull_max_count_once", 5,
		"specify max count of pulling operation logs once")
	opLogPullInterval := flag.Uint64("oplog_pull_interval", 10,
		"specify interval of pulling operation logs in second")
	opLogPullTimeout := flag.Uint64("oplog_pull_timeout", 30,
		"specify timeout of pulling operation logs in second")

	host := flag.String("host", "localhost", "specify listen host")
	certFile := flag.String("certfile", "", "specify cert file, "+
		"downgrade HTTPS(10443) to HTTP(10080) if it is set empty or inexistent file")
	keyFile := flag.String("keyfile", "", "specify key file, "+
		"downgrade HTTPS(10443) to HTTP(10080) if it is set empty or inexistent file")
	stage := flag.String("stage", "debug", "sepcify runtime stage (debug, test, prod)")
	configHome := flag.String("config", CONFIG_HOME_DIR, "sepcify config home path")
	logHome := flag.String("log", LOG_HOME_DIR, "specify log home path")
	cpuProfileFile := flag.String("cpuprofile", "", "specify cpu profile output file, "+
		"cpu profiling will be fully disabled if not provided")
	memProfileFile := flag.String("memprofile", "", "specify heap dump file, "+
		"memory profiling will be fully disabled if not provided")

	flag.Parse()

	ClusterGroup = *clusterGroup
	MemberMode = *memberMode
	OPLogMaxSeqGapToPull = *opLogMaxSeqGapToPull
	OPLogPullMaxCountOnce = *opLogPullMaxCountOnce
	OPLogPullInterval = time.Duration(*opLogPullInterval) * time.Second
	OPLogPullTimeout = time.Duration(*opLogPullTimeout) * time.Second

	Host = *host
	CertFile = *certFile
	KeyFile = *keyFile
	Stage = *stage
	ConfigHome = *configHome
	LogHome = *logHome
	CpuProfileFile = *cpuProfileFile
	MemProfileFile = *memProfileFile
}
