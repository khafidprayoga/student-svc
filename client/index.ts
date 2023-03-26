import { StudentServiceClient } from './gen/student/v2/student.client';
import { GenderType, StudentNationality } from './gen/student/v2/student';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';
import { ChannelCredentials } from '@grpc/grpc-js';
import { Timestamp } from './gen/google/protobuf/timestamp';

const transport = new GrpcTransport({
  host: 'localhost:5080',
  channelCredentials: ChannelCredentials.createInsecure(),
});
const client = new StudentServiceClient(transport);

const [, , studentId] = process.argv;
(async () => {
  try {
    const { response: detail } = await client.getDetailStudent({
      studentId,
    });
    var pronouns = '';

    switch (detail.gender) {
      case GenderType.FEMALE:
        pronouns = 'she/her';
        break;
      case GenderType.MALE:
        pronouns = 'he/him';
        break;
    }

    let dob, created_at, updated_at;

    if (detail.birthDate) {
      dob = Timestamp.toDate(detail.birthDate);
    }

    if (detail.createdAt) {
      created_at = Timestamp.toDate(detail.createdAt);
    }

    if (detail.updatedAt) {
      updated_at = Timestamp.toDate(detail.updatedAt);
    }
    console.log(`
    Hello, i am ${detail.fullName},
    pronouns ${pronouns},
    My birth date ${dob}
    My Address ${detail.address}
    My favorite activies ${detail.hobbies}
    I live as ${StudentNationality[detail.nationality].toLocaleLowerCase()}
    Regards, ${detail.email}
    ==========
    Track Id: ${detail.id}
    User Created At: ${created_at}
    Last Updated At: ${updated_at} 
    `);
  } catch (error) {
    console.error(error);
  }
})();
