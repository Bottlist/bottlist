import { CenterLayout } from '../../components/CenterLayout';
import { TextField } from '../../components/TextField';
import { yupResolver } from '@hookform/resolvers/yup';
import { Box, Button, Stack, Typography } from '@mui/material';
import { useForm, FormProvider } from 'react-hook-form';
import * as yup from 'yup';

const schema = yup
  .object({
    name: yup.string().required(),
    furigana: yup.string().required(),
    owner: yup
      .object({
        name: yup.string().required(),
        furigana: yup.string().required(),
      })
      .required(),
    email: yup.string().required(),
    landline: yup.string().required(),
    mobile: yup.string().required(),
    dueDate: yup.string().required(),
    businessHours: yup.string().required(),
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

  const submit = () => {
    console.log('TBD');
  };

  return (
    <CenterLayout spacing={2}>
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(submit)}>
          <Stack spacing={1} paddingY={2} alignItems="center">
            <Typography variant="h5">新規登録</Typography>
            <TextField placeholder="店名" _key="name" />
            <TextField placeholder="ふりがな" _key="furigana" />
            <TextField placeholder="代表者名" _key="owner.name" />
            <TextField placeholder="ふりがな" _key="owner.furigana" />
            <TextField placeholder="メールアドレス" _key="email" />
            <TextField placeholder="TEL（店）" _key="landline" />
            <TextField placeholder="TEL（携帯）" _key="mobile" />
            <TextField placeholder="ボトルキープ期限" _key="dueDate" />
            <TextField placeholder="営業日時" _key=" businessHours" />
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
