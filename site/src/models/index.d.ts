import { ModelInit, MutableModel, PersistentModelConstructor } from "@aws-amplify/datastore";





export declare class UserProfile {
  readonly id: string;
  readonly awsConfigs?: string;
  constructor(init: ModelInit<UserProfile>);
  static copyOf(source: UserProfile, mutator: (draft: MutableModel<UserProfile>) => MutableModel<UserProfile> | void): UserProfile;
}