import {
  Avatar,
  Box,
  Button,
  Container,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableCellProps,
  TableContainer,
  TableRow,
  Typography,
} from '@mui/material';
import { request } from '../../utils/axiosUtils';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { useState } from 'react';
import { Navigation } from './components/Navigation';
import { useDropzone } from 'react-dropzone';
import { UpperLeftLogo } from '../../components/logo/UpperLeftLogo';
import { PasswordResetButton } from '../../components/button/PasswordResetButton';
import { Modal } from '../../components/Modal';

const Td = (props: TableCellProps) => (
  <TableCell sx={{ borderColor: '#555555' }} {...props} />
);
const Th = (props: TableCellProps) => (
  <TableCell {...props} component="th" scope="row" sx={{ border: 'none' }} />
);

export const Profile = () => {
  const queryClient = useQueryClient();
  const { data } = useQuery({
    queryKey: ['me'],
    queryFn: () =>
      request({
        url: '/me',
        method: 'get',
      }).then((r) => r.data),
  });
  const [isConfirmModalOpen, setIsConfirmModalOpen] = useState(false);
  const [isDoneModalOpen, setIsDoneModalOpen] = useState(false);
  const { getRootProps, getInputProps } = useDropzone({
    accept: { 'image/png': ['.png'] },
    onDrop: (acceptedFiles) => {
      const fr = new FileReader();
      fr.onload = () => queryClient.invalidateQueries(['me']);
      fr.readAsDataURL(acceptedFiles[0]);
    },
  });
  if (!data) return null;
  return (
    <>
      <UpperLeftLogo />
      <Container>
        <Stack spacing={3} marginTop="93px">
          <Box {...getRootProps()}>
            <input {...getInputProps()} />
            <Avatar
              alt={data?.screen_name}
              src={data?.img}
              sx={{ width: 180, height: 180, marginX: 'auto' }}
            />
          </Box>
          <Container>
            <TableContainer
              sx={{
                borderTop: '1px solid #555555',
                borderBottom: '1px solid #555555',
              }}
            >
              <Table>
                <TableBody>
                  <TableRow>
                    <Th>ふりがな</Th>
                    <Td>{data?.first_name_huri}</Td>
                    <Td>{data?.last_name_huri}</Td>
                  </TableRow>
                  <TableRow>
                    <Th>名前</Th>
                    <Td>{data?.first_name}</Td>
                    <Td>{data?.last_name}</Td>
                  </TableRow>
                  <TableRow>
                    <Th>
                      ボトルネーム
                      <br />
                      ニックネーム
                    </Th>
                    <Td colSpan={2}>{data?.screen_name}</Td>
                  </TableRow>
                  <TableRow>
                    <Th>生年月日</Th>
                    <Td colSpan={2}>{data?.birthday}</Td>
                  </TableRow>
                  <TableRow>
                    <Th>メールアドレス</Th>
                    <Td sx={{}} colSpan={2}>
                      {data?.email}
                    </Td>
                  </TableRow>
                </TableBody>
              </Table>
            </TableContainer>
          </Container>
          <PasswordResetButton onClick={() => setIsConfirmModalOpen(true)} />
        </Stack>
      </Container>
      <Navigation />
      <Modal
        open={isConfirmModalOpen}
        onClose={() => setIsConfirmModalOpen(false)}
      >
        <Stack
          textAlign="center"
          spacing={5}
          paddingY="2rem"
          alignItems="center"
        >
          <Typography>
            下記の新しいメールアドレスに
            <br />
            情報を送信します
          </Typography>
          <Typography
            sx={{ backgroundColor: '#EEEEEE' }}
            width="fit-content"
            paddingX="40px"
            paddingY="16px"
          >
            {data?.email}
          </Typography>
          <Stack direction="row" justifyContent="space-evenly" width="100%">
            <Button
              onClick={() =>
                request({
                  url: '/auth/password/reset/link',
                  method: 'post',
                  data: { email: data?.email },
                }).finally(
                  () => (setIsConfirmModalOpen(false), setIsDoneModalOpen(true))
                )
              }
            >
              はい
            </Button>
            <Button
              onClick={() => setIsConfirmModalOpen(false)}
              color="secondary"
            >
              いいえ
            </Button>
          </Stack>
        </Stack>
      </Modal>
      <Modal open={isDoneModalOpen} onClose={() => setIsDoneModalOpen(false)}>
        <Stack spacing={5}>
          <Typography textAlign="center">メールを送信しました</Typography>
          <Box textAlign="end">
            <Button color="secondary" onClick={() => setIsDoneModalOpen(false)}>
              戻る
            </Button>
          </Box>
        </Stack>
      </Modal>
    </>
  );
};
