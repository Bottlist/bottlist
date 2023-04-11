import { DatePicker } from '../../components/DatePicker';
import { Logo } from '../../components/Logo';
import { TextField } from '../../components/TextField';
import { yupResolver } from '@hookform/resolvers/yup';
import { Box, Button, Container, MenuItem, Paper, Stack } from '@mui/material';
import { useForm, FormProvider } from 'react-hook-form';
import * as yup from 'yup';

const schema = yup
  .object({
    date: yup.date().required(),
    shop: yup.string().required(),
    type: yup.string().required(),
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
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(submit)}>
        <Container>
          <Stack spacing={2}>
            <Logo />
            <Paper>
              <Box minHeight={256} minWidth={256}>
                ボトル画像選択
              </Box>
            </Paper>
            <DatePicker label="開封日入力" control={control} _key="date" />
            <TextField placeholder="店検索" _key="shop" />
            <TextField label="種類選択" _key="type" select>
              <MenuItem>日本酒</MenuItem>
              <MenuItem>カクテル</MenuItem>
            </TextField>
            <TextField placeholder="銘柄入力" _key="type" />
            <Box textAlign="center">
              <Button color="secondary" type="submit">
                登録
              </Button>
            </Box>
          </Stack>
        </Container>
      </form>
    </FormProvider>
  );
};
