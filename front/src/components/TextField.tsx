import { TextFieldProps, TextField as MuiTextField } from '@mui/material';
import { FieldValues, Path, useFormContext } from 'react-hook-form';

export const TextField = <T extends FieldValues>(
  props: TextFieldProps & { _key: Path<T> }
) => {
  const {
    register,
    getValues,
    formState: { errors },
  } = useFormContext<T>();
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
