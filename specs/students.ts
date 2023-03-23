const student = {
  id: '497f6eca-6276-4993-bfeb-53cbbbba6f08', // random uuid v4
  full_name: 'Jane Doe', // only a-zA-Z
  birth_date: '2019-08-24T14:15:22Z',
  gender: 'male', // male female other
  address: 'Jalan Jend. Sudirman No. 321',
  hobbies: ['football', 'fishing'],
  email: 'janedoe@example.com',
  nationality: 'foreign', // citizen and foreign
};

console.log(JSON.stringify(student, null, '  '));
