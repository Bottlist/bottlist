import {
  Avatar,
  Box,
  Button,
  Modal,
  Paper,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableRow,
  Typography,
} from '@mui/material';
import { Logo } from '../../components/Logo';
import { request } from '../../utils/axiosUtils';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { useState } from 'react';
import { Navigation } from '../../components/Navigation';
import { useDropzone } from 'react-dropzone';

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
  const [open, setOpen] = useState(false);
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
      <Stack>
        <Logo />
        <Box {...getRootProps()}>
          <input {...getInputProps()} />
          <Avatar alt={data?.nickname} src={data?.img} />
        </Box>
        <TableContainer component={Paper}>
          <Table>
            <TableBody>
              <TableRow>
                <TableCell component="th" scope="row">
                  ふりがな
                </TableCell>
                <TableCell>{data?.family_name_kana}</TableCell>
                <TableCell>{data?.last_name_kana}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell component="th" scope="row">
                  名前
                </TableCell>
                <TableCell>{data?.family_name}</TableCell>
                <TableCell>{data?.last_name}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell component="th" scope="row">
                  ボトルネーム
                  <br />
                  ニックネーム
                </TableCell>
                <TableCell colSpan={2}>{data?.nickname}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell component="th" scope="row">
                  生年月日
                </TableCell>
                <TableCell colSpan={2}>{data?.birthdate}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell component="th" scope="row">
                  メールアドレス
                </TableCell>
                <TableCell colSpan={2}>{data?.email}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
        <Button onClick={() => setOpen(true)}>
          パスワードの変更はこちらから
        </Button>
        <Navigation />
      </Stack>
      <Modal open={open} onClose={() => setOpen(false)}>
        <Paper>
          <Typography>下記の新しいメールアドレスに情報を送信します</Typography>
          <Typography>{data?.email}</Typography>
          <Button
            onClick={() =>
              request({
                url: '/password/reset/link',
                method: 'post',
                data: { email: data?.email },
              }).finally(() => setOpen(false))
            }
          >
            はい
          </Button>
          <Button onClick={() => setOpen(false)}>いいえ</Button>
        </Paper>
      </Modal>
    </>
  );
};
