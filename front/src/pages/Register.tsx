import { Box, Button, Stack, Typography } from '@mui/material';

import { CenterLayout } from '../components/CenterLayout';
import * as yup from 'yup';
import { useForm, FormProvider } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { TextField } from '../components/TextField';
import { DatePicker } from '../components/DatePicker';

const schema = yup
  .object({
    name: yup.string().required(),
    furigana: yup.string().required(),
    nickname: yup.string().required(),
    birthdate: yup.string().required(),
    email: yup.string().required(),
    password: yup.string().required(),
    passwordConfirmation: yup
      .string()
      .required()
      .oneOf([yup.ref('password')], 'Passwords must match'),
  })
  .required();
type FormType = yup.InferType<typeof schema>;

export const Register = () => {
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });
  const { control } = methods;

  const submit = () => {
    console.log('TBD');
  };

  return (
    <CenterLayout spacing={2}>
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(submit)}>
          <Stack spacing={1} paddingY={2} alignItems="center">
            <Typography variant="h5">新規登録</Typography>
            <TextField placeholder="名前" _key="name" />
            <TextField placeholder="ふりがな" _key="furigana" />
            <TextField
              placeholder="ボトルネーム・ニックネーム"
              _key="nickname"
            />
            <DatePicker label="生年月日" _key="birthdate" control={control} />
            <TextField placeholder="メールアドレス" _key="email" />
            <TextField placeholder="パスワード" _key="password" />
            <TextField
              placeholder="パスワード（確認）"
              _key="passwordConfirmation"
            />
            <Box paddingTop={4} width="100%">
              <Button fullWidth type="submit">
                登録
              </Button>
            </Box>
          </Stack>
        </form>
      </FormProvider>
    </CenterLayout>
  );
};
