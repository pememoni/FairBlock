// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRevealDecryption } from "./types/fairblock/tx";
import { MsgSubmitShare } from "./types/fairblock/tx";
import { MsgSubmitTarget } from "./types/fairblock/tx";
import { MsgSubmitEncrypted } from "./types/fairblock/tx";
import { MsgCommitDecryption } from "./types/fairblock/tx";


const types = [
  ["/pememoni.fairblock.fairblock.MsgRevealDecryption", MsgRevealDecryption],
  ["/pememoni.fairblock.fairblock.MsgSubmitShare", MsgSubmitShare],
  ["/pememoni.fairblock.fairblock.MsgSubmitTarget", MsgSubmitTarget],
  ["/pememoni.fairblock.fairblock.MsgSubmitEncrypted", MsgSubmitEncrypted],
  ["/pememoni.fairblock.fairblock.MsgCommitDecryption", MsgCommitDecryption],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgRevealDecryption: (data: MsgRevealDecryption): EncodeObject => ({ typeUrl: "/pememoni.fairblock.fairblock.MsgRevealDecryption", value: MsgRevealDecryption.fromPartial( data ) }),
    msgSubmitShare: (data: MsgSubmitShare): EncodeObject => ({ typeUrl: "/pememoni.fairblock.fairblock.MsgSubmitShare", value: MsgSubmitShare.fromPartial( data ) }),
    msgSubmitTarget: (data: MsgSubmitTarget): EncodeObject => ({ typeUrl: "/pememoni.fairblock.fairblock.MsgSubmitTarget", value: MsgSubmitTarget.fromPartial( data ) }),
    msgSubmitEncrypted: (data: MsgSubmitEncrypted): EncodeObject => ({ typeUrl: "/pememoni.fairblock.fairblock.MsgSubmitEncrypted", value: MsgSubmitEncrypted.fromPartial( data ) }),
    msgCommitDecryption: (data: MsgCommitDecryption): EncodeObject => ({ typeUrl: "/pememoni.fairblock.fairblock.MsgCommitDecryption", value: MsgCommitDecryption.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
