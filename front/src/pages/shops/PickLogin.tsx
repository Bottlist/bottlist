import { ChoicesLayout } from '../../components/ChoicesLayout';

export function PickLogin() {
  return (
    <ChoicesLayout
      items={[
        { text: 'ログイン', link: '/shops/login' },
        {
          text: '新規登録',
          link: '/shops/register',
        },
      ]}
    />
  );
}
