// @ts-check
import { initSchema } from '@aws-amplify/datastore';
import { schema } from './schema';



const { UserProfile } = initSchema(schema);

export {
  UserProfile
};