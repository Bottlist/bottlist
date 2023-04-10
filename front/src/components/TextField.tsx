import { TextFieldProps, TextField as MuiTextField } from '@mui/material';
import { useFormContext } from 'react-hook-form';

export const TextField = (props: TextFieldProps & { _key: string }) => {
  const {
    register,
    getValues,
    formState: { errors },
  } = useFormContext();
  const { _key } = props;

  return (
    <MuiTextField
      {...props}
      {...register(_key)}
      defaultValue={getValues(_key)}
      error={!!errors[_key]}
      helperText={errors[_key]?.message?.toString()}
    />
  );
};
