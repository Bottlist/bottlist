import { Modal, Paper, Typography } from '@mui/material';
import { request } from '../../utils/axiosUtils';
import * as yup from 'yup';
import { Button } from '../../components/Button';
import { GrayTextField } from '../../components/GrayTextField';
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
    <Modal open onClose={props.onClose}>
      <Paper>
        <FormProvider {...methods}>
          <form onSubmit={onSubmit}>
            <Typography>パスワードをお忘れの方</Typography>
            <TextField
              _key="email"
              label="ご登録中のメールアドレス"
              alignLabel="center"
            />
            <Button type="submit">送信</Button>
          </form>
        </FormProvider>
      </Paper>
    </Modal>
  );
};
