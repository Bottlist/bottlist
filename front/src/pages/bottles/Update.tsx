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
    type: yup.string().required(),
  })
  .required();
type FormType = yup.InferType<typeof schema>;

export const Update = () => {
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
            <TextField label="種類選択" _key="type" select>
              <MenuItem>日本酒</MenuItem>
              <MenuItem>カクテル</MenuItem>
            </TextField>
            <TextField placeholder="銘柄入力" _key="type" />
            <Box textAlign="center">
              <Button color="secondary">更新</Button>
            </Box>
            <Box textAlign="center">
              <Button color="secondary">削除</Button>
            </Box>
          </Stack>
        </Container>
      </form>
    </FormProvider>
  );
};
