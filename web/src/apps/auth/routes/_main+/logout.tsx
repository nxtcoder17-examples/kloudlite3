import { useNavigate } from '@remix-run/react';
import { BrandLogo } from '@kloudlite/design-system/branding/brand-logo';
import { handleError, sleep } from '~/root/lib/utils/common';
import { useAuthApi } from '~/auth/server/gql/api-provider';
import { toast } from 'react-toastify';
import useDebounce from '~/root/lib/client/hooks/use-debounce';

const LogoutPage = () => {
  const navigate = useNavigate();
  const api = useAuthApi();

  useDebounce(
    () => {
      (async () => {
        try {
          const { errors } = await api.logout({});
          if (errors) {
            throw errors[0];
          }

          toast.warn('Logged out successfully');
          await sleep(1000);
          navigate('/login');
        } catch (error) {
          handleError(error);
        }
      })();
    },
    1000,
    []
  );
  return (
    <div className="flex flex-col items-center justify-center gap-3xl h-full">
      <BrandLogo detailed={false} size={56} />
      <span className="headingLg text-text-strong">Logging out...</span>
    </div>
  );
};

export default LogoutPage;
