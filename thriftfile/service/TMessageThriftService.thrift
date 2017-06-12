namespace go thrift.source.message.service
namespace java com.loopin.service.thrift.protocal.message.service

include "../struct/request/TMessageReq.thrift"

service TMessageThriftService {

    void sendMessage(1:TMessageReq.TSendMessageCondition sendMessageCondition);

}