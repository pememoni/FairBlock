/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "pememoni.fairblock.fairblock";
const baseMsgSubmitEncrypted = {
    creator: "",
    messageHash: "",
    encryption: "",
    targetHeight: "",
    deposit: "",
};
export const MsgSubmitEncrypted = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.messageHash !== "") {
            writer.uint32(18).string(message.messageHash);
        }
        if (message.encryption !== "") {
            writer.uint32(26).string(message.encryption);
        }
        if (message.targetHeight !== "") {
            writer.uint32(34).string(message.targetHeight);
        }
        if (message.deposit !== "") {
            writer.uint32(42).string(message.deposit);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSubmitEncrypted };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.messageHash = reader.string();
                    break;
                case 3:
                    message.encryption = reader.string();
                    break;
                case 4:
                    message.targetHeight = reader.string();
                    break;
                case 5:
                    message.deposit = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSubmitEncrypted };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.messageHash !== undefined && object.messageHash !== null) {
            message.messageHash = String(object.messageHash);
        }
        else {
            message.messageHash = "";
        }
        if (object.encryption !== undefined && object.encryption !== null) {
            message.encryption = String(object.encryption);
        }
        else {
            message.encryption = "";
        }
        if (object.targetHeight !== undefined && object.targetHeight !== null) {
            message.targetHeight = String(object.targetHeight);
        }
        else {
            message.targetHeight = "";
        }
        if (object.deposit !== undefined && object.deposit !== null) {
            message.deposit = String(object.deposit);
        }
        else {
            message.deposit = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.messageHash !== undefined &&
            (obj.messageHash = message.messageHash);
        message.encryption !== undefined && (obj.encryption = message.encryption);
        message.targetHeight !== undefined &&
            (obj.targetHeight = message.targetHeight);
        message.deposit !== undefined && (obj.deposit = message.deposit);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSubmitEncrypted };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.messageHash !== undefined && object.messageHash !== null) {
            message.messageHash = object.messageHash;
        }
        else {
            message.messageHash = "";
        }
        if (object.encryption !== undefined && object.encryption !== null) {
            message.encryption = object.encryption;
        }
        else {
            message.encryption = "";
        }
        if (object.targetHeight !== undefined && object.targetHeight !== null) {
            message.targetHeight = object.targetHeight;
        }
        else {
            message.targetHeight = "";
        }
        if (object.deposit !== undefined && object.deposit !== null) {
            message.deposit = object.deposit;
        }
        else {
            message.deposit = "";
        }
        return message;
    },
};
const baseMsgSubmitEncryptedResponse = {};
export const MsgSubmitEncryptedResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSubmitEncryptedResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = {
            ...baseMsgSubmitEncryptedResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgSubmitEncryptedResponse,
        };
        return message;
    },
};
const baseMsgCommitDecryption = {
    creator: "",
    plaintextHash: "",
    plaintextDecryptorHash: "",
};
export const MsgCommitDecryption = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.plaintextHash !== "") {
            writer.uint32(18).string(message.plaintextHash);
        }
        if (message.plaintextDecryptorHash !== "") {
            writer.uint32(26).string(message.plaintextDecryptorHash);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCommitDecryption };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.plaintextHash = reader.string();
                    break;
                case 3:
                    message.plaintextDecryptorHash = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCommitDecryption };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.plaintextHash !== undefined && object.plaintextHash !== null) {
            message.plaintextHash = String(object.plaintextHash);
        }
        else {
            message.plaintextHash = "";
        }
        if (object.plaintextDecryptorHash !== undefined &&
            object.plaintextDecryptorHash !== null) {
            message.plaintextDecryptorHash = String(object.plaintextDecryptorHash);
        }
        else {
            message.plaintextDecryptorHash = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.plaintextHash !== undefined &&
            (obj.plaintextHash = message.plaintextHash);
        message.plaintextDecryptorHash !== undefined &&
            (obj.plaintextDecryptorHash = message.plaintextDecryptorHash);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCommitDecryption };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.plaintextHash !== undefined && object.plaintextHash !== null) {
            message.plaintextHash = object.plaintextHash;
        }
        else {
            message.plaintextHash = "";
        }
        if (object.plaintextDecryptorHash !== undefined &&
            object.plaintextDecryptorHash !== null) {
            message.plaintextDecryptorHash = object.plaintextDecryptorHash;
        }
        else {
            message.plaintextDecryptorHash = "";
        }
        return message;
    },
};
const baseMsgCommitDecryptionResponse = {};
export const MsgCommitDecryptionResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCommitDecryptionResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = {
            ...baseMsgCommitDecryptionResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCommitDecryptionResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    SubmitEncrypted(request) {
        const data = MsgSubmitEncrypted.encode(request).finish();
        const promise = this.rpc.request("pememoni.fairblock.fairblock.Msg", "SubmitEncrypted", data);
        return promise.then((data) => MsgSubmitEncryptedResponse.decode(new Reader(data)));
    }
    CommitDecryption(request) {
        const data = MsgCommitDecryption.encode(request).finish();
        const promise = this.rpc.request("pememoni.fairblock.fairblock.Msg", "CommitDecryption", data);
        return promise.then((data) => MsgCommitDecryptionResponse.decode(new Reader(data)));
    }
}
