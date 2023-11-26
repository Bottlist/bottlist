import { Stack, Typography } from '@mui/material';
import { Logo } from '../../components/logo/Logo';
import { Navigation } from './components/Navigation';
import { TextField as UntypedTextField } from '../../components/TextField';
import * as yup from 'yup';
import { FormProvider, useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';

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
const TextField = UntypedTextField<FormType>;

export const Settings = () => {
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });

  const submit = () => {
    console.log('TBD');
  };

  return (
    <>
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(submit)}>
          <Stack>
            <Logo />
            <Typography>各種設定</Typography>
            <Stack direction="row">
              <Typography>ふりがな</Typography>
              <TextField _key="furigana" />
            </Stack>
            <Stack direction="row">
              <Typography>店名</Typography>
              <TextField _key="name" />
            </Stack>
            <Stack direction="row">
              <Typography>ふりがな</Typography>
              <TextField _key="owner.furigana" />
            </Stack>
            <Stack direction="row">
              <Typography>代表者名</Typography>
              <TextField _key="owner.name" />
            </Stack>
            <Stack direction="row">
              <Typography>メールアドレス</Typography>
              <TextField _key="email" />
            </Stack>
            <Stack direction="row">
              <Typography>電話番号</Typography>
              <TextField _key="mobile" />
            </Stack>
            <Stack direction="row">
              <Typography>キープ期限</Typography>
              {/* <DatePicker label="生年月日" _key="" control={control} /> */}
            </Stack>
            <Stack direction="row">
              <Typography>営業時間</Typography>
              {/* <TextField _key="furigana" /> */}
            </Stack>
          </Stack>
        </form>
      </FormProvider>
      <Navigation />
    </>
  );
};
