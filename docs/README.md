# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [common.proto](#common-proto)
    - [IPAddress](#grpc_api-IPAddress)
    - [IPv4Address](#grpc_api-IPv4Address)
    - [IPv6Address](#grpc_api-IPv6Address)
    - [JSON](#grpc_api-JSON)
    - [Port](#grpc_api-Port)
    - [Timestamp](#grpc_api-Timestamp)
  
- [storage.proto](#storage-proto)
    - [DatedBotReply](#grpc_api-DatedBotReply)
    - [DatedEdge](#grpc_api-DatedEdge)
    - [DatedFailedTry](#grpc_api-DatedFailedTry)
    - [DisconnectRequest](#grpc_api-DisconnectRequest)
    - [DisconnectResponse](#grpc_api-DisconnectResponse)
    - [RegisterSessionRequest](#grpc_api-RegisterSessionRequest)
    - [RegisterSessionResponse](#grpc_api-RegisterSessionResponse)
    - [SessionToken](#grpc_api-SessionToken)
    - [StoreDatedBotReplyBatchRequest](#grpc_api-StoreDatedBotReplyBatchRequest)
    - [StoreDatedBotReplyBatchResponse](#grpc_api-StoreDatedBotReplyBatchResponse)
    - [StoreDatedEdgeBatchRequest](#grpc_api-StoreDatedEdgeBatchRequest)
    - [StoreDatedEdgeBatchResponse](#grpc_api-StoreDatedEdgeBatchResponse)
    - [StoreDatedFailedTryBatchRequest](#grpc_api-StoreDatedFailedTryBatchRequest)
    - [StoreDatedFailedTryBatchResponse](#grpc_api-StoreDatedFailedTryBatchResponse)
  
    - [DisconnectReason](#grpc_api-DisconnectReason)
  
    - [BMSStorageService](#grpc_api-BMSStorageService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common.proto



<a name="grpc_api-IPAddress"></a>

### IPAddress
A type for carrying an IPv4 or an IPv6 address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| v4 | [IPv4Address](#grpc_api-IPv4Address) |  |  |
| v6 | [IPv6Address](#grpc_api-IPv6Address) |  |  |






<a name="grpc_api-IPv4Address"></a>

### IPv4Address
A type for carrying an IPv4 address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [fixed32](#fixed32) |  | Actual value (required). |






<a name="grpc_api-IPv6Address"></a>

### IPv6Address
A type for carrying an IPv6 address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | Actual value (required). |






<a name="grpc_api-JSON"></a>

### JSON
A type for carrying a JSON string.

Actually this type has no validator, but the server still checks if the provided string is valid JSON.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json_string | [string](#string) |  | Actual value (required). |






<a name="grpc_api-Port"></a>

### Port
A type for carrying a valid port number.

Note: Excludes port 0, because it cannot be used to communicate externally.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [uint32](#uint32) |  | Actual value (required). |






<a name="grpc_api-Timestamp"></a>

### Timestamp
A type for carrying a timestamp with nano second resolution.

Note: Although this has a nano second resolution, the BMS database stores time only in micro second resolution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unix_seconds | [int64](#int64) |  | Seconds since the Unix epoch (1970-01-01) (required). |
| unix_nano_seconds | [int64](#int64) |  | Nano seconds to add to the given unix seconds (optional). |





 

 

 

 



<a name="storage-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage.proto



<a name="grpc_api-DatedBotReply"></a>

### DatedBotReply
The type which holds exactly one bot reply.

This message will be just used to fill batches which in turn will be sent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [Timestamp](#grpc_api-Timestamp) |  | The timestamp the bot reply was observed (required). |
| bot_id | [string](#string) | optional | The bot&#39;s ID, if any in the respective botnet (optional). |
| ip | [IPAddress](#grpc_api-IPAddress) |  | The IP address of the bot (required). |
| port | [Port](#grpc_api-Port) |  | The port of the bot (required). |
| other_data | [JSON](#grpc_api-JSON) | optional | Potential other interesting data observed with the bot reply (optional). |






<a name="grpc_api-DatedEdge"></a>

### DatedEdge
The type which holds exactly one edge.

This message will be just used to fill batches which in turn will be sent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [Timestamp](#grpc_api-Timestamp) |  | The timestamp the edge was observed (required). |
| src_bot_id | [string](#string) | optional | The source bot&#39;s ID, if any in the respective botnet (optional). |
| src_ip | [IPAddress](#grpc_api-IPAddress) |  | The IP address of the source bot (required). |
| src_port | [Port](#grpc_api-Port) |  | The port of the source bot (required). |
| dst_bot_id | [string](#string) | optional | The destination bot&#39;s ID, if any in the respective botnet (optional). |
| dst_ip | [IPAddress](#grpc_api-IPAddress) |  | The IP address of the destination bot (required). |
| dst_port | [Port](#grpc_api-Port) |  | The port of the destination bot (required). |






<a name="grpc_api-DatedFailedTry"></a>

### DatedFailedTry
The type which holds exactly one failed bot contact try.

This message will be just used to fill batches which in turn will be sent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [Timestamp](#grpc_api-Timestamp) |  | The timestamp the (failed) bot contact try was attempted (required). |
| bot_id | [string](#string) | optional | The bot&#39;s ID, if any in the respective botnet (optional). |
| ip | [IPAddress](#grpc_api-IPAddress) |  | The IP address of the bot (required). |
| port | [Port](#grpc_api-Port) |  | The port of the bot (required). |
| reason | [string](#string) | optional | The the reason the contacted bot could not be reached (optional). |
| other_data | [JSON](#grpc_api-JSON) | optional | Potential other interesting data observed with the (failed) bot contact try (optional). |






<a name="grpc_api-DisconnectRequest"></a>

### DisconnectRequest
The request the client sends when initiating a disconnect (which in turn explicitly ends the session).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reason | [DisconnectReason](#grpc_api-DisconnectReason) |  | The reason the clients wants to end the session due (optional). |






<a name="grpc_api-DisconnectResponse"></a>

### DisconnectResponse
The response the server sends when receiving a disconnect request.






<a name="grpc_api-RegisterSessionRequest"></a>

### RegisterSessionRequest
The request used for registering a monitoring session.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| botnet_id | [string](#string) |  | The botnet id the current monitoring session is monitoring (required). |
| campaign_id | [string](#string) | optional | The campaign id the current monitoring session should be part of (optional). |
| frequency | [uint32](#uint32) | optional | The used frequency for crawling/populating for the current monitoring session (optional). |
| public_ip | [IPAddress](#grpc_api-IPAddress) | optional | The public IP address the monitor uses for the current monitoring session (optional). |
| monitor_port | [Port](#grpc_api-Port) | optional | The port the monitoring session uses to listen for other bots (optional). |
| config_json | [JSON](#grpc_api-JSON) | optional | Potential further configuration data for the current monitoring session (optional). |






<a name="grpc_api-RegisterSessionResponse"></a>

### RegisterSessionResponse
The response the server sends when receiving a session registration request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_token | [SessionToken](#grpc_api-SessionToken) |  | The 32bit session token for all further requests of this session (required). |
| session_timeout | [uint32](#uint32) |  | The timeout in seconds after which the server automatically ends the session (required).

Precisely means: The session expires if the time since last message from the client is more than the timeout. |
| last_inserted_bot_reply | [Timestamp](#grpc_api-Timestamp) | optional | The timestamp of the last successfully inserted bot reply (optional).

Not set when there is no last entry (or the server might not want to add it for other reasons). |
| last_inserted_failed_try | [Timestamp](#grpc_api-Timestamp) | optional | The timestamp of the last successfully inserted failed try (optional).

Not set when there is no last entry (or the server might not want to add it for other reasons). |
| last_inserted_edge | [Timestamp](#grpc_api-Timestamp) | optional | The timestamp of the last successfully inserted edge (optional).

Not set when there is no last entry (or the server might not want to add it for other reasons). |






<a name="grpc_api-SessionToken"></a>

### SessionToken
A type for carrying the 32bit session token.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [bytes](#bytes) |  | Actual value (required). |






<a name="grpc_api-StoreDatedBotReplyBatchRequest"></a>

### StoreDatedBotReplyBatchRequest
A batch of dated bot replies.

This is the message which actually will be streamed to the server.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_id | [uint32](#uint32) |  | The ID to identify this batch (required).

This ID will be used by the server to later reference this batch.

It is the responsibility of the client to use this ID wisely, the server will not verify uniqueness and will just use the ID it got from the client to reference batches. It is recommended to simply use sequentially incrementing numbers per session. |
| replies | [DatedBotReply](#grpc_api-DatedBotReply) | repeated | The bot replies to include in this batch (required). |






<a name="grpc_api-StoreDatedBotReplyBatchResponse"></a>

### StoreDatedBotReplyBatchResponse
The response the server sends when receiving a dated bot reply batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_id | [uint32](#uint32) |  | The ID of the batch which was successfully received (required).

The referenced batch can now be removed from the clients cache. |






<a name="grpc_api-StoreDatedEdgeBatchRequest"></a>

### StoreDatedEdgeBatchRequest
A batch of dated edges.

This is the message which actually will be streamed to the server.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_id | [uint32](#uint32) |  | The ID to identify this batch (required).

This ID will be used by the server to later reference this batch.

It is the responsibility of the client to use this ID wisely, the server will not verify uniqueness and will just use the ID it got from the client to reference batches. It is recommended to simply use sequentially incrementing numbers per session. |
| edges | [DatedEdge](#grpc_api-DatedEdge) | repeated | The edges to include in this batch (required). |






<a name="grpc_api-StoreDatedEdgeBatchResponse"></a>

### StoreDatedEdgeBatchResponse
The response the server sends when receiving a dated edge batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_id | [uint32](#uint32) |  | The ID of the batch which was successfully received (required).

The referenced batch can now be removed from the clients cache. |






<a name="grpc_api-StoreDatedFailedTryBatchRequest"></a>

### StoreDatedFailedTryBatchRequest
A batch of dated (failed) bot contact tries.

This is the message which actually will be streamed to the server.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_id | [uint32](#uint32) |  | The ID to identify this batch (required).

This ID will be used by the server to later reference this batch.

It is the responsibility of the client to use this ID wisely, the server will not verify uniqueness and will just use the ID it got from the client to reference batches. It is recommended to simply use sequentially incrementing numbers per session. |
| tries | [DatedFailedTry](#grpc_api-DatedFailedTry) | repeated | The (failed) bot contact tries to include in this batch (required). |






<a name="grpc_api-StoreDatedFailedTryBatchResponse"></a>

### StoreDatedFailedTryBatchResponse
The response the server sends when receiving a dated (failed) bot contact try batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_id | [uint32](#uint32) |  | The ID of the batch which was successfully received (required).

The referenced batch can now be removed from the clients cache. |





 


<a name="grpc_api-DisconnectReason"></a>

### DisconnectReason
Possible reasons a client wants to end the session with.

| Name | Number | Description |
| ---- | ------ | ----------- |
| DISCONNECT_REASON_UNSPECIFIED | 0 | The client wants to specify no reason to end the session. |
| DISCONNECT_REASON_FINISHED | 1 | The client has done its purpose and thereby wants to end the session. |
| DISCONNECT_REASON_BE_RIGHT_BACK | 2 | The clients wants to end the session in order to reconnect soon. |
| DISCONNECT_REASON_BE_RIGHT_BACK_WITH_NEW_CONFIG | 3 | The clients wants to end the session in order to reconnect soon with a new configuration. |
| DISCONNECT_REASON_CLIENT_ERROR | 4 | The client had some error and thereby wants to the the session. |
| DISCONNECT_REASON_OTHER | 5 | The client wants to end the session because of some other reason. |


 

 


<a name="grpc_api-BMSStorageService"></a>

### BMSStorageService
The BMS Storage service definition.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterSession | [RegisterSessionRequest](#grpc_api-RegisterSessionRequest) | [RegisterSessionResponse](#grpc_api-RegisterSessionResponse) | The method to register a new monitoring session.

This method gives you back a session token which is needed for all further requests corresponding to the session. Furthermore it (optionally) can tell the client the time of the last bot reply batch, failed try batch and edge batch for the last session of the same monitor (and thereby giving an indicator for the last successful insert).

Note: This method needs authentication via providing the 256bit auth token as `auth-token-bin` and the monitor ID as `auth-monitor-id` in the metadata (last given value wins). |
| StoreDatedBotReplyBatch | [StoreDatedBotReplyBatchRequest](#grpc_api-StoreDatedBotReplyBatchRequest) stream | [StoreDatedBotReplyBatchResponse](#grpc_api-StoreDatedBotReplyBatchResponse) stream | The method to send a stream of dated bot replies.

The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`. The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).

Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).

@todo: Document errors which can occur |
| StoreDatedFailedTryBatch | [StoreDatedFailedTryBatchRequest](#grpc_api-StoreDatedFailedTryBatchRequest) stream | [StoreDatedFailedTryBatchResponse](#grpc_api-StoreDatedFailedTryBatchResponse) stream | The method to send a stream of dated (failed) bot contact tries.

The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`. The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).

Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).

@todo: Document errors which can occur |
| StoreDatedEdgeBatch | [StoreDatedEdgeBatchRequest](#grpc_api-StoreDatedEdgeBatchRequest) stream | [StoreDatedEdgeBatchResponse](#grpc_api-StoreDatedEdgeBatchResponse) stream | The method to send a stream of dated bot edges.

The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`. The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).

Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins). Note: The server will not accept batches with more than 10000 bot replies.

@todo: Document errors which can occur |
| Disconnect | [DisconnectRequest](#grpc_api-DisconnectRequest) | [DisconnectResponse](#grpc_api-DisconnectResponse) | The method to disconnect the current monitoring session.

Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins). Note: The server will not accept batches with more than 10000 edges.

@todo: Document errors which can occur |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

