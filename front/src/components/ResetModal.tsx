import { Container, Modal, Paper, Stack, Typography } from '@mui/material';
import { request } from '../utils/axiosUtils';
import * as yup from 'yup';
import { Button } from './Button';
import { GrayTextField } from './GrayTextField';
import { FormProvider, useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { useNavigate } from 'react-router';

const schema = yup
  .object({
    email: yup.string().email().required(),
  })
  .required();
type FormType = yup.InferType<typeof schema>;

const TextField = GrayTextField<FormType>;

export const ResetModal = (props: { onClose: () => void }) => {
  const navigate = useNavigate();
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });

  const onSubmit = methods.handleSubmit(({ email }) =>
    request({
      url: '/auth/password/reset/link',
      method: 'post',
      data: { email },
    }).then(() => navigate('/top'))
  );

  return (
    <Modal open onClose={props.onClose} sx={{ top: '50%' }}>
      <Container sx={{ transform: 'translate(0%, -50%)' }}>
        <Paper>
          <FormProvider {...methods}>
            <form onSubmit={onSubmit}>
              <Stack textAlign="center" spacing={5} paddingY="2rem">
                <Typography>パスワードをお忘れの方</Typography>
                <TextField
                  _key="email"
                  label="ご登録中のメールアドレス"
                  alignLabel="center"
                />
                <Stack direction="row" justifyContent="space-evenly">
                  <Button type="submit" width="fit-content">
                    送信
                  </Button>
                  <Button type="submit" color="secondary" width="fit-content">
                    戻る
                  </Button>
                </Stack>
              </Stack>
            </form>
          </FormProvider>
        </Paper>
      </Container>
    </Modal>
  );
};
