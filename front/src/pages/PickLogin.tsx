import { Choices } from '../components/Choices';

export function PickLogin() {
  return (
    <Choices
      items={[
        { text: 'ログイン' },
        {
          text: '新規登録',
        },
      ]}
    />
  );
}
