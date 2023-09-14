import { ChoicesLayout } from '../../components/ChoicesLayout';

export function PickLogin() {
  return (
    <ChoicesLayout
      items={[
        { text: 'ログイン', link: '/login' },
        {
          text: '新規登録',
          link: '/users/register',
        },
      ]}
    />
  );
}
