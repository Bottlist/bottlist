import { Choices } from '../components/Choices';

export function PickLogin() {
  return (
    <Choices
      items={[
        { text: 'ログイン', link: '/login' },
        {
          text: '新規登録',
        },
      ]}
    />
  );
}
