import { Navigation } from './components/Navigation';
import { TextField as UntypedTextField } from '../../components/TextField';
import * as yup from 'yup';
import { FormProvider, useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { UpperLeftLogo } from '../../components/logo/UpperLeftLogo';
import { Container, Stack, Typography } from '@mui/material';
import { ReactNode } from 'react';

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

const Td = (props: { children: ReactNode }) => (
  <td style={{ borderTop: '1px solid', borderBottom: '1px solid' }}>
    {props.children}
  </td>
);

export const Settings = () => {
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });

  const submit = () => {
    console.log('TBD');
  };

  return (
    <>
      <UpperLeftLogo />
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(submit)}>
          <Container sx={{ paddingTop: '70px' }}>
            <Typography variant="h5" textAlign="center">
              各種設定
            </Typography>
            <table
              style={{
                borderCollapse: 'collapse',
                borderTop: '1px solid',
                borderBottom: '1px solid',
                marginTop: '20px',
              }}
            >
              <tr>
                <th>ふりがな</th>
                <Td>
                  <TextField _key="furigana" />
                </Td>
              </tr>
              <tr>
                <th>店名</th>
                <Td>
                  <TextField _key="name" />
                </Td>
              </tr>
              <tr>
                <th>ふりがな</th>
                <Td>
                  <TextField _key="owner.furigana" />
                </Td>
              </tr>
              <tr>
                <th>代表者名</th>
                <Td>
                  <TextField _key="owner.name" />
                </Td>
              </tr>
              <tr>
                <th>メールアドレス</th>
                <Td>
                  <TextField _key="email" />
                </Td>
              </tr>
              <tr>
                <th>電話番号</th>
                <Td>
                  <TextField _key="mobile" />
                </Td>
              </tr>
              <tr>
                <th>キープ期限</th>
                <Td>
                  <Stack
                    direction="row"
                    alignItems="center"
                    justifyContent="space-around"
                  >
                    <Typography fontSize={12}>キープ開始日から</Typography>
                    <TextField _key="mobile" sx={{ width: '70px' }} />
                    <Typography fontSize={12}> ヶ月</Typography>
                  </Stack>
                </Td>
              </tr>
              <tr>
                <th>営業時間</th>
                <Td>{/* {} */}</Td>
              </tr>
            </table>
          </Container>
        </form>
      </FormProvider>
      <Navigation />
    </>
  );
};
