import { promisify } from 'util'
import { credentials } from '@grpc/grpc-js'
import { UserManagerClient } from '../proto/users_grpc_pb.js'
import { CreateUserReply, CreateUserRequest, SearchUserReply, SearchUserRequest, UserType } from '../proto/users_pb.js'

void main()
async function main () {
  const client = new UserManagerClient(
    'localhost:8080',
    credentials.createInsecure()
  )
  
  const createUserRequest =
    new CreateUserRequest()
      .setName('Minhyeok Park')
      .setAge(19)
      .setType(UserType.UT_CTO)
  
  const createUser = promisify(client.createUser.bind(client))

  const createUserReply =
    await createUser(createUserRequest) as CreateUserReply

  console.log('User created:', createUserReply.getId())

  const searchUserRequest =
    new SearchUserRequest()

  const searchUser = promisify(client.searchUser.bind(client))

  const searchUserReply =
    await searchUser(searchUserRequest) as SearchUserReply

  console.log(searchUserReply.toObject())
}
