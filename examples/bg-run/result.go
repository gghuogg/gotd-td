package main


/*

GOROOT=G:\golang\go1.20.5 #gosetup
GOPATH=G:\golang\go1.20.5\bin;G:\golang\go1.17.5.windows-amd64\go\bin;G:\golang\golang\go\bin;G:\golang\go1.18.5.windows-amd64\go\bin #gosetup
G:\golang\go1.20.5\bin\go.exe build -o C:\Users\Administrator\AppData\Local\Temp\___go_build_github_com_gotd_td_examples_bg_run.exe github.com/gotd/td/examples/bg-run #gosetup
C:\Users\Administrator\AppData\Local\Temp\___go_build_github_com_gotd_td_examples_bg_run.exe #gosetup
29453132 60e35bbbefb9d5cb6b395f1d04b1d588
2023-10-26T21:47:10.547+0800	INFO	conn.mtproto.rpc	rpc/engine.go:42	Initialized	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "retry_interval": "5s", "max_retries": 5}
2023-10-26T21:47:10.555+0800	INFO	telegram/connect.go:114	Starting	{"v": "v0.83.0"}
2023-10-26T21:47:10.555+0800	DEBUG	conn.mtproto	mtproto/conn.go:198	Run: start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:10.557+0800	INFO	conn.mtproto	mtproto/connect.go:32	Generating new auth key	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:10.557+0800	DEBUG	conn.mtproto	mtproto/connect.go:65	Initializing new key exchange	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "timeout": "1m0s", "context_deadline": "2023-10-26T21:47:45.555+0800"}
2023-10-26T21:47:10.557+0800	DEBUG	conn.mtproto.exchange	exchange/client_flow.go:26	Sending ReqPqMultiRequest	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:11.254+0800	DEBUG	conn.mtproto.exchange	exchange/client_flow.go:37	Received server ResPQ	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:11.340+0800	DEBUG	conn.mtproto.exchange	exchange/client_flow.go:76	PQ decomposing complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "took": "86.0172ms"}
2023-10-26T21:47:11.340+0800	DEBUG	conn.mtproto.exchange	exchange/client_flow.go:120	Sending ReqDHParamsRequest	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:12.128+0800	DEBUG	conn.mtproto.exchange	exchange/client_flow.go:129	Received server ServerDHParams	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:12.466+0800	DEBUG	conn.mtproto.exchange	exchange/client_flow.go:212	Sending SetClientDHParamsRequest	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.exchange	exchange/client_flow.go:224	Received server DhGen	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:12.995+0800	INFO	conn.mtproto	mtproto/connect.go:38	Auth key generated	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "2.4388526s"}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "readLoop"}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "pingLoop"}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "saltsLoop"}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.read	mtproto/read.go:135	Read loop started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "handleClose"}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "ackLoop"}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "userCallback"}
2023-10-26T21:47:12.995+0800	DEBUG	conn	manager/conn.go:191	Initializing	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263356315935772}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263356315935772}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294263356315935772}
2023-10-26T21:47:12.995+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294263356315935772}
2023-10-26T21:47:13.287+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x73f1f8dc", "type_name": "message_container", "size_bytes": 88, "msg_id": 7294263348915631105}
2023-10-26T21:47:13.287+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x9ec20908", "type_name": "new_session_created#9ec20908", "size_bytes": 28, "msg_id": 7294263348915631105}
2023-10-26T21:47:13.287+0800	DEBUG	conn.mtproto	mtproto/handle_session_created.go:21	Session created	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "unique_id": 6241328006955329244, "first_msg_id": 7294263356315935772, "first_msg_time": "2023-10-26T21:47:12.995+0800"}
2023-10-26T21:47:13.287+0800	INFO	conn	manager/conn.go:79	SessionInit	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:13.287+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xb921bd04", "type_name": "get_future_salts#b921bd04", "msg_id": 7294263359902907268}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 400, "msg_id": 7294263349314465793}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x3072cfa1", "type_name": "gzip", "size_bytes": 388, "msg_id": 7294263356315935772}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:39	Decompressed	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xcc1a241e", "type_name": "config#cc1a241e", "size_bytes": 960, "msg_id": 7294263356315935772}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263356315935772}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto	mtproto/rpc.go:47	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263356315935772}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "userCallback", "elapsed": "402.5779ms"}
2023-10-26T21:47:13.398+0800	DEBUG	telegram/connect.go:84	Ready	{"v": "v0.83.0"}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263360013480968}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263360013480968}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294263360013480968}
2023-10-26T21:47:13.398+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294263360013480968}
2023-10-26T21:47:13.399+0800	DEBUG	telegram/session.go:89	Data saved	{"v": "v0.83.0", "key_id": "6915ab732ac4f553"}
2023-10-26T21:47:13.399+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x62d6b459", "type_name": "msgs_ack#62d6b459", "size_bytes": 20, "msg_id": 7294263348915631105}
2023-10-26T21:47:13.399+0800	DEBUG	conn.mtproto	mtproto/handle_ack.go:17	Received ack	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_ids": [7294263356315935772]}
2023-10-26T21:47:13.399+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:15	Acknowledge callback not set	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263356315935772}
2023-10-26T21:47:13.635+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xae500895", "type_name": "future_salts#ae500895", "size_bytes": 84, "msg_id": 7294263350325613569}
2023-10-26T21:47:13.635+0800	DEBUG	conn.mtproto	mtproto/handle_future_salts.go:23	Got future salts	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "server_time": "2023-10-26T21:47:10.000+0800"}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294263351126231041}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294263360013480968}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263360013480968, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263360013480968}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294263360013480968}
2023-10-26T21:47:13.814+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "416.2216ms", "ongoing": 0}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto	mtproto/conn.go:172	Closing	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:13.814+0800	DEBUG	telegram/connect.go:155	Callback returned, stopping	{"v": "v0.83.0"}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "saltsLoop", "error": "context canceled", "elapsed": "818.7995ms"}
2023-10-26T21:47:13.814+0800	INFO	conn.mtproto.rpc	rpc/engine.go:271	Close called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "pingLoop", "error": "ping loop: context canceled", "errorVerbose": "ping loop:\n    github.com/gotd/td/internal/mtproto.(*Conn).pingLoop\n        G:/goData/telegram/td-main/internal/mtproto/ping.go:107\n  - context canceled", "elapsed": "818.7995ms"}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "ackLoop", "error": "acl: context canceled", "errorVerbose": "acl:\n    github.com/gotd/td/internal/mtproto.(*Conn).ackLoop\n        G:/goData/telegram/td-main/internal/mtproto/ack.go:33\n  - context canceled", "elapsed": "818.7995ms"}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "handleClose", "elapsed": "818.7995ms"}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto.read	mtproto/read.go:141	Read loop done	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "reason": "context canceled"}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "readLoop", "error": "context canceled", "elapsed": "818.7995ms"}
2023-10-26T21:47:13.814+0800	DEBUG	conn.mtproto	mtproto/conn.go:214	Run: end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-26T21:47:13.814+0800	INFO	telegram/connect.go:164	Closed	{"v": "v0.83.0"}

进程完成，并显示退出代码 0
















GOROOT=G:\golang\go1.20.5 #gosetup
GOPATH=G:\golang\go1.20.5\bin;G:\golang\go1.17.5.windows-amd64\go\bin;G:\golang\golang\go\bin;G:\golang\go1.18.5.windows-amd64\go\bin #gosetup
G:\golang\go1.20.5\bin\go.exe build -o C:\Users\Administrator\AppData\Local\Temp\___go_build_github_com_gotd_td_examples_bg_run.exe github.com/gotd/td/examples/bg-run #gosetup
C:\Users\Administrator\AppData\Local\Temp\___go_build_github_com_gotd_td_examples_bg_run.exe #gosetup
29453132 60e35bbbefb9d5cb6b395f1d04b1d588
2023-10-27T08:58:47.020+0800	INFO	conn.mtproto.rpc	rpc/engine.go:42	Initialized	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "retry_interval": "5s", "max_retries": 5}
2023-10-27T08:58:47.026+0800	INFO	telegram/connect.go:114	Starting	{"v": "v0.83.0"}
2023-10-27T08:58:47.027+0800	INFO	telegram/session.go:46	Connection restored from state	{"v": "v0.83.0", "addr": "", "key_id": "6915ab732ac4f553"}
2023-10-27T08:58:47.027+0800	INFO	conn.mtproto.rpc	rpc/engine.go:42	Initialized	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "retry_interval": "5s", "max_retries": 5}
2023-10-27T08:58:47.027+0800	DEBUG	conn.mtproto	mtproto/conn.go:198	Run: start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:47.029+0800	INFO	conn.mtproto	mtproto/connect.go:44	Key already exists	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto	mtproto/connect.go:49	Generating new session id	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "pingLoop"}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "readLoop"}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "saltsLoop"}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "handleClose"}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "userCallback"}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.read	mtproto/read.go:135	Read loop started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "ackLoop"}
2023-10-27T08:58:47.029+0800	DEBUG	conn	manager/conn.go:191	Initializing	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294436421056233892}
2023-10-27T08:58:47.029+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294436421056233892}
2023-10-27T08:58:47.497+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xedab447b", "type_name": "bad_server_salt#edab447b", "size_bytes": 28, "msg_id": 7294436406555659265}
2023-10-27T08:58:47.498+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.498+0800	DEBUG	conn.mtproto	mtproto/rpc.go:36	Setting server salt	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:47.498+0800	INFO	conn.mtproto	mtproto/salt.go:21	Salt updated	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "old": 61391858328394797, "new": -2877645450721313719}
2023-10-27T08:58:47.498+0800	INFO	conn.mtproto	mtproto/rpc.go:41	Retrying request after basMsgErr	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.498+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.498+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294436421056233892}
2023-10-27T08:58:47.498+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294436421056233892}
2023-10-27T08:58:47.708+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x73f1f8dc", "type_name": "message_container", "size_bytes": 88, "msg_id": 7294436407503127557}
2023-10-27T08:58:47.708+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x9ec20908", "type_name": "new_session_created#9ec20908", "size_bytes": 28, "msg_id": 7294436407503127557}
2023-10-27T08:58:47.708+0800	DEBUG	conn.mtproto	mtproto/handle_session_created.go:21	Session created	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "unique_id": 9013209988380027768, "first_msg_id": 7294436421056233892, "first_msg_time": "2023-10-27T08:58:47.029+0800"}
2023-10-27T08:58:47.708+0800	INFO	conn	manager/conn.go:79	SessionInit	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:47.708+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xb921bd04", "type_name": "get_future_salts#b921bd04", "msg_id": 7294436421735810992}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 400, "msg_id": 7294436407521484801}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x3072cfa1", "type_name": "gzip", "size_bytes": 388, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:39	Decompressed	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xcc1a241e", "type_name": "config#cc1a241e", "size_bytes": 960, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto	mtproto/rpc.go:42	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "userCallback", "elapsed": "684.8435ms"}
2023-10-27T08:58:47.713+0800	DEBUG	telegram/connect.go:84	Ready	{"v": "v0.83.0"}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421741077392}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421741077392}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294436421741077392}
2023-10-27T08:58:47.713+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294436421741077392}
2023-10-27T08:58:47.714+0800	DEBUG	telegram/session.go:89	Data saved	{"v": "v0.83.0", "key_id": "6915ab732ac4f553"}
2023-10-27T08:58:47.714+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x62d6b459", "type_name": "msgs_ack#62d6b459", "size_bytes": 20, "msg_id": 7294436407503127557}
2023-10-27T08:58:47.714+0800	DEBUG	conn.mtproto	mtproto/handle_ack.go:17	Received ack	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_ids": [7294436421056233892]}
2023-10-27T08:58:47.714+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:15	Acknowledge callback not set	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421056233892}
2023-10-27T08:58:47.919+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xae500895", "type_name": "future_salts#ae500895", "size_bytes": 84, "msg_id": 7294436408410984449}
2023-10-27T08:58:47.919+0800	DEBUG	conn.mtproto	mtproto/handle_future_salts.go:23	Got future salts	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "server_time": "2023-10-27T08:58:44.000+0800"}
2023-10-27T08:58:47.927+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294436408444090369}
2023-10-27T08:58:47.927+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294436421741077392}
2023-10-27T08:58:47.927+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421741077392, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-27T08:58:47.927+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421741077392}
2023-10-27T08:58:47.927+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436421741077392}
2023-10-27T08:58:47.927+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "213.6562ms", "ongoing": 0}
2023-10-27T08:58:47.927+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x7abe77ec", "type_name": "ping#7abe77ec", "msg_id": 7294436421954733592}
2023-10-27T08:58:48.137+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x347773c5", "type_name": "pong#347773c5", "size_bytes": 20, "msg_id": 7294436409349181441}
2023-10-27T08:58:48.137+0800	DEBUG	conn.mtproto	mtproto/ping.go:43	Pong	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:48.137+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425459537788}
2023-10-27T08:58:48.137+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425459537788}
2023-10-27T08:58:48.137+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294436425459537788}
2023-10-27T08:58:48.137+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294436425459537788}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294436410256824321}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294436425459537788}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425459537788, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425459537788}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425459537788}
2023-10-27T08:58:48.350+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "213.1778ms", "ongoing": 0}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425672715588}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425672715588}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294436425672715588}
2023-10-27T08:58:48.350+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294436425672715588}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294436411174135809}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294436425672715588}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425672715588, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425672715588}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294436425672715588}
2023-10-27T08:58:48.563+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "212.9495ms", "ongoing": 0}
2023-10-27T08:58:48.563+0800	DEBUG	telegram/connect.go:155	Callback returned, stopping	{"v": "v0.83.0"}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "saltsLoop", "error": "context canceled", "elapsed": "1.5344639s"}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto	mtproto/conn.go:172	Closing	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:48.563+0800	INFO	conn.mtproto.rpc	rpc/engine.go:271	Close called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "pingLoop", "error": "ping loop: context canceled", "errorVerbose": "ping loop:\n    github.com/gotd/td/internal/mtproto.(*Conn).pingLoop\n        G:/goData/telegram/td-main/internal/mtproto/ping.go:107\n  - context canceled", "elapsed": "1.5344639s"}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "ackLoop", "error": "acl: context canceled", "errorVerbose": "acl:\n    github.com/gotd/td/internal/mtproto.(*Conn).ackLoop\n        G:/goData/telegram/td-main/internal/mtproto/ack.go:33\n  - context canceled", "elapsed": "1.5344639s"}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "handleClose", "elapsed": "1.5344639s"}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto.read	mtproto/read.go:141	Read loop done	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "reason": "context canceled"}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "readLoop", "error": "context canceled", "elapsed": "1.5344639s"}
2023-10-27T08:58:48.563+0800	DEBUG	conn.mtproto	mtproto/conn.go:214	Run: end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T08:58:48.563+0800	INFO	telegram/connect.go:164	Closed	{"v": "v0.83.0"}

进程完成，并显示退出代码 0


















29453132 60e35bbbefb9d5cb6b395f1d04b1d588
2023-10-27T09:08:24.126+0800	INFO	conn.mtproto.rpc	rpc/engine.go:42	Initialized	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "retry_interval": "5s", "max_retries": 5}
2023-10-27T09:08:24.132+0800	INFO	telegram/connect.go:114	Starting	{"v": "v0.83.0"}
2023-10-27T09:08:24.133+0800	INFO	telegram/session.go:46	Connection restored from state	{"v": "v0.83.0", "addr": "", "key_id": "6915ab732ac4f553"}
2023-10-27T09:08:24.133+0800	INFO	conn.mtproto.rpc	rpc/engine.go:42	Initialized	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "retry_interval": "5s", "max_retries": 5}
2023-10-27T09:08:24.133+0800	DEBUG	conn.mtproto	mtproto/conn.go:198	Run: start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:24.134+0800	INFO	conn.mtproto	mtproto/connect.go:44	Key already exists	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto	mtproto/connect.go:49	Generating new session id	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "readLoop"}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "saltsLoop"}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.read	mtproto/read.go:135	Read loop started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "pingLoop"}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "userCallback"}
2023-10-27T09:08:24.134+0800	DEBUG	conn	manager/conn.go:191	Initializing	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "handleClose"}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "ackLoop"}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899358021584}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899358021584}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294438899358021584}
2023-10-27T09:08:24.134+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294438899358021584}
2023-10-27T09:08:24.581+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x73f1f8dc", "type_name": "message_container", "size_bytes": 88, "msg_id": 7294438885083829249}
2023-10-27T09:08:24.581+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x9ec20908", "type_name": "new_session_created#9ec20908", "size_bytes": 28, "msg_id": 7294438885083829249}
2023-10-27T09:08:24.581+0800	DEBUG	conn.mtproto	mtproto/handle_session_created.go:21	Session created	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "unique_id": 821417370499213528, "first_msg_id": 7294438899358021584, "first_msg_time": "2023-10-27T09:08:24.134+0800"}
2023-10-27T09:08:24.581+0800	INFO	conn	manager/conn.go:79	SessionInit	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:24.581+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xb921bd04", "type_name": "get_future_salts#b921bd04", "msg_id": 7294438899805031184}
2023-10-27T09:08:24.586+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 400, "msg_id": 7294438885103221761}
2023-10-27T09:08:24.586+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x3072cfa1", "type_name": "gzip", "size_bytes": 388, "msg_id": 7294438899358021584}
2023-10-27T09:08:24.586+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:39	Decompressed	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xcc1a241e", "type_name": "config#cc1a241e", "size_bytes": 960, "msg_id": 7294438899358021584}
2023-10-27T09:08:24.586+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899358021584}
2023-10-27T09:08:24.586+0800	DEBUG	conn.mtproto	mtproto/rpc.go:47	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899358021584}
2023-10-27T09:08:24.586+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "userCallback", "elapsed": "452.2247ms"}
2023-10-27T09:08:24.586+0800	DEBUG	telegram/connect.go:84	Ready	{"v": "v0.83.0"}
2023-10-27T09:08:24.587+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899810795284}
2023-10-27T09:08:24.587+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899810795284}
2023-10-27T09:08:24.587+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294438899810795284}
2023-10-27T09:08:24.587+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294438899810795284}
2023-10-27T09:08:24.587+0800	DEBUG	telegram/session.go:89	Data saved	{"v": "v0.83.0", "key_id": "6915ab732ac4f553"}
2023-10-27T09:08:24.587+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x62d6b459", "type_name": "msgs_ack#62d6b459", "size_bytes": 20, "msg_id": 7294438885083829249}
2023-10-27T09:08:24.587+0800	DEBUG	conn.mtproto	mtproto/handle_ack.go:17	Received ack	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_ids": [7294438899358021584]}
2023-10-27T09:08:24.587+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:15	Acknowledge callback not set	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899358021584}
2023-10-27T09:08:24.796+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xae500895", "type_name": "future_salts#ae500895", "size_bytes": 84, "msg_id": 7294438886006491137}
2023-10-27T09:08:24.796+0800	DEBUG	conn.mtproto	mtproto/handle_future_salts.go:23	Got future salts	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "server_time": "2023-10-27T09:08:20.000+0800"}
2023-10-27T09:08:24.799+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294438886021498881}
2023-10-27T09:08:24.799+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294438899810795284}
2023-10-27T09:08:24.799+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899810795284, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-27T09:08:24.799+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899810795284}
2023-10-27T09:08:24.799+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438899810795284}
2023-10-27T09:08:24.800+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "212.7993ms", "ongoing": 0}
2023-10-27T09:08:24.800+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x7abe77ec", "type_name": "ping#7abe77ec", "msg_id": 7294438900023604184}
ping
phoneGetCallConfig
2023-10-27T09:08:25.006+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x347773c5", "type_name": "pong#347773c5", "size_bytes": 20, "msg_id": 7294438886909972481}
2023-10-27T09:08:25.006+0800	DEBUG	conn.mtproto	mtproto/ping.go:43	Pong	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:25.006+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903524849580}
2023-10-27T09:08:25.006+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903524849580}
2023-10-27T09:08:25.006+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294438903524849580}
2023-10-27T09:08:25.006+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294438903524849580}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294438887828210689}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294438903524849580}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903524849580, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903524849580}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903524849580}
2023-10-27T09:08:25.224+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "218.3801ms", "ongoing": 0}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903743229680}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903743229680}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294438903743229680}
2023-10-27T09:08:25.224+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294438903743229680}
AccountCheckUsername
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294438888897654785}
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294438903743229680}
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903743229680, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903743229680}
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294438903743229680}
2023-10-27T09:08:25.469+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "244.5505ms", "ongoing": 0}
2023-10-27T09:08:25.469+0800	DEBUG	telegram/connect.go:155	Callback returned, stopping	{"v": "v0.83.0"}
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto	mtproto/conn.go:172	Closing	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "saltsLoop", "error": "context canceled", "elapsed": "1.3347913s"}
2023-10-27T09:08:25.469+0800	INFO	conn.mtproto.rpc	rpc/engine.go:271	Close called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "pingLoop", "error": "ping loop: context canceled", "errorVerbose": "ping loop:\n    github.com/gotd/td/internal/mtproto.(*Conn).pingLoop\n        G:/goData/telegram/td-main/internal/mtproto/ping.go:107\n  - context canceled", "elapsed": "1.3347913s"}
2023-10-27T09:08:25.469+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "ackLoop", "error": "acl: context canceled", "errorVerbose": "acl:\n    github.com/gotd/td/internal/mtproto.(*Conn).ackLoop\n        G:/goData/telegram/td-main/internal/mtproto/ack.go:33\n  - context canceled", "elapsed": "1.3347913s"}
2023-10-27T09:08:25.470+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "handleClose", "elapsed": "1.3357909s"}
2023-10-27T09:08:25.470+0800	DEBUG	conn.mtproto.read	mtproto/read.go:141	Read loop done	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "reason": "context canceled"}
2023-10-27T09:08:25.470+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "readLoop", "error": "context canceled", "elapsed": "1.3357909s"}
2023-10-27T09:08:25.470+0800	DEBUG	conn.mtproto	mtproto/conn.go:214	Run: end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:08:25.470+0800	INFO	telegram/connect.go:164	Closed	{"v": "v0.83.0"}




















GOROOT=G:\golang\go1.20.5 #gosetup
GOPATH=G:\golang\go1.20.5\bin;G:\golang\go1.17.5.windows-amd64\go\bin;G:\golang\golang\go\bin;G:\golang\go1.18.5.windows-amd64\go\bin #gosetup
G:\golang\go1.20.5\bin\go.exe build -o C:\Users\Administrator\AppData\Local\Temp\___go_build_github_com_gotd_td_examples_bg_run.exe github.com/gotd/td/examples/bg-run #gosetup
C:\Users\Administrator\AppData\Local\Temp\___go_build_github_com_gotd_td_examples_bg_run.exe #gosetup
29453132 60e35bbbefb9d5cb6b395f1d04b1d588
bg.Connect
2023-10-27T09:46:09.150+0800	INFO	conn.mtproto.rpc	rpc/engine.go:42	Initialized	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "retry_interval": "5s", "max_retries": 5}
2023-10-27T09:46:09.156+0800	INFO	telegram/connect.go:114	Starting	{"v": "v0.83.0"}
2023-10-27T09:46:09.157+0800	INFO	telegram/session.go:46	Connection restored from state	{"v": "v0.83.0", "addr": "", "key_id": "6915ab732ac4f553"}
2023-10-27T09:46:09.157+0800	INFO	conn.mtproto.rpc	rpc/engine.go:42	Initialized	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "retry_interval": "5s", "max_retries": 5}
2023-10-27T09:46:09.157+0800	DEBUG	conn.mtproto	mtproto/conn.go:198	Run: start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:09.159+0800	INFO	conn.mtproto	mtproto/connect.go:44	Key already exists	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto	mtproto/connect.go:49	Generating new session id	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "ackLoop"}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "saltsLoop"}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "pingLoop"}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "userCallback"}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "readLoop"}
2023-10-27T09:46:09.159+0800	DEBUG	conn	manager/conn.go:191	Initializing	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.read	mtproto/read.go:135	Read loop started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627483666324}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627483666324}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294448627483666324}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294448627483666324}
2023-10-27T09:46:09.159+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task started	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "handleClose"}
2023-10-27T09:46:09.645+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x73f1f8dc", "type_name": "message_container", "size_bytes": 88, "msg_id": 7294448613124337665}
2023-10-27T09:46:09.645+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x9ec20908", "type_name": "new_session_created#9ec20908", "size_bytes": 28, "msg_id": 7294448613124337665}
2023-10-27T09:46:09.645+0800	DEBUG	conn.mtproto	mtproto/handle_session_created.go:21	Session created	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "unique_id": 3276433044071181183, "first_msg_id": 7294448627483666324, "first_msg_time": "2023-10-27T09:46:09.159+0800"}
2023-10-27T09:46:09.645+0800	INFO	conn	manager/conn.go:79	SessionInit	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:09.645+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xb921bd04", "type_name": "get_future_salts#b921bd04", "msg_id": 7294448627969568424}
2023-10-27T09:46:09.649+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 400, "msg_id": 7294448613143742465}
2023-10-27T09:46:09.649+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x3072cfa1", "type_name": "gzip", "size_bytes": 388, "msg_id": 7294448627483666324}
2023-10-27T09:46:09.649+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:39	Decompressed	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xcc1a241e", "type_name": "config#cc1a241e", "size_bytes": 960, "msg_id": 7294448627483666324}
2023-10-27T09:46:09.649+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627483666324}
2023-10-27T09:46:09.649+0800	DEBUG	conn.mtproto	mtproto/rpc.go:47	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627483666324}
2023-10-27T09:46:09.649+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "userCallback", "elapsed": "490.5948ms"}
2023-10-27T09:46:09.649+0800	DEBUG	telegram/connect.go:84	Ready	{"v": "v0.83.0"}
2023-10-27T09:46:09.650+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627974773624}
2023-10-27T09:46:09.650+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627974773624}
2023-10-27T09:46:09.650+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294448627974773624}
2023-10-27T09:46:09.650+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294448627974773624}
2023-10-27T09:46:09.651+0800	DEBUG	telegram/session.go:89	Data saved	{"v": "v0.83.0", "key_id": "6915ab732ac4f553"}
2023-10-27T09:46:09.651+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x62d6b459", "type_name": "msgs_ack#62d6b459", "size_bytes": 20, "msg_id": 7294448613124337665}
2023-10-27T09:46:09.651+0800	DEBUG	conn.mtproto	mtproto/handle_ack.go:17	Received ack	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_ids": [7294448627483666324]}
2023-10-27T09:46:09.651+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:15	Acknowledge callback not set	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627483666324}
2023-10-27T09:46:09.864+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xae500895", "type_name": "future_salts#ae500895", "size_bytes": 84, "msg_id": 7294448614070671361}
2023-10-27T09:46:09.864+0800	DEBUG	conn.mtproto	mtproto/handle_future_salts.go:23	Got future salts	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "server_time": "2023-10-27T09:46:05.000+0800"}
2023-10-27T09:46:09.941+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294448614393663489}
2023-10-27T09:46:09.941+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294448627974773624}
2023-10-27T09:46:09.941+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627974773624, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-27T09:46:09.941+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627974773624}
2023-10-27T09:46:09.941+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448627974773624}
2023-10-27T09:46:09.941+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "290.6708ms", "ongoing": 0}
2023-10-27T09:46:09.941+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x7abe77ec", "type_name": "ping#7abe77ec", "msg_id": 7294448628265444424}
ping
AccountCheckUsername
2023-10-27T09:46:10.162+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x347773c5", "type_name": "pong#347773c5", "size_bytes": 20, "msg_id": 7294448615344535553}
2023-10-27T09:46:10.162+0800	DEBUG	conn.mtproto	mtproto/ping.go:43	Pong	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:10.162+0800	DEBUG	conn.mtproto	mtproto/rpc.go:29	Invoke start	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448631781450120}
2023-10-27T09:46:10.162+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:87	Do called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448631781450120}
2023-10-27T09:46:10.162+0800	DEBUG	conn.mtproto.rpc	rpc/ack.go:34	Waiting for acknowledge	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "ack_id": 7294448631781450120}
2023-10-27T09:46:10.162+0800	DEBUG	conn.mtproto	mtproto/new_encrypted_msg.go:62	Request	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xbf9459b7", "type_name": "invokeWithoutUpdates#bf9459b7", "msg_id": 7294448631781450120}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto	mtproto/handle_message.go:19	Handle message	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0xf35c6d01", "type_name": "rpc_result", "size_bytes": 44, "msg_id": 7294448616305685505}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:24	Handle result	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "type_id": "0x2144ca19", "type_name": "rpc_error#2144ca19", "size_bytes": 32, "msg_id": 7294448631781450120}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto	mtproto/handle_result.go:53	Got error	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448631781450120, "err_code": 401, "err_msg": "AUTH_KEY_UNREGISTERED"}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto.rpc	rpc/engine.go:99	Handler called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448631781450120}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto	mtproto/rpc.go:44	Invoke end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "msg_id": 7294448631781450120}
2023-10-27T09:46:10.385+0800	DEBUG	conn	manager/conn.go:113	Invoke	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "duration": "223.5027ms", "ongoing": 0}
2023-10-27T09:46:10.385+0800	DEBUG	telegram/connect.go:155	Callback returned, stopping	{"v": "v0.83.0"}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto	mtproto/conn.go:172	Closing	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "saltsLoop", "error": "context canceled", "elapsed": "1.2263192s"}
2023-10-27T09:46:10.385+0800	INFO	conn.mtproto.rpc	rpc/engine.go:271	Close called	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "pingLoop", "error": "ping loop: context canceled", "errorVerbose": "ping loop:\n    github.com/gotd/td/internal/mtproto.(*Conn).pingLoop\n        G:/goData/telegram/td-main/internal/mtproto/ping.go:107\n  - context canceled", "elapsed": "1.2263192s"}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "ackLoop", "error": "acl: context canceled", "errorVerbose": "acl:\n    github.com/gotd/td/internal/mtproto.(*Conn).ackLoop\n        G:/goData/telegram/td-main/internal/mtproto/ack.go:33\n  - context canceled", "elapsed": "1.2263192s"}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task complete	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "handleClose", "elapsed": "1.2263192s"}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto.read	mtproto/read.go:141	Read loop done	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "reason": "context canceled"}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto.group	tdsync/cancel_group.go:48	Task stopped	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2, "task": "readLoop", "error": "context canceled", "elapsed": "1.2263192s"}
2023-10-27T09:46:10.385+0800	DEBUG	conn.mtproto	mtproto/conn.go:214	Run: end	{"v": "v0.83.0", "conn_id": 0, "dc_id": 2}
2023-10-27T09:46:10.385+0800	INFO	telegram/connect.go:164	Closed	{"v": "v0.83.0"}
boolerr rpcDoRequest: rpc error code 401: AUTH_KEY_UNREGISTERED
false

进程完成，并显示退出代码 0









*/












