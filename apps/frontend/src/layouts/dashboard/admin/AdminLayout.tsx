import CoffeeLoadingAnimation from '@/components/shared/CoffeeLoadingAnimation';
import { useAdminDashboard } from '@/features/dashboard/admin/hooks/admin.hook';
import HeaderAdmin from '@/features/dashboard/admin/components/header-components/sections/Header/AdminHeader';
import { Outlet } from 'react-router-dom';
import CoffeeErrorAnimation from '@/components/shared/CoffeeErrorAnimation';

const AdminLayout = () => {
  const { isLoading, isError, error } = useAdminDashboard(true);

  return (
    <div className="flex min-h-screen flex-col bg-[#f8f3e9]">
      <HeaderAdmin />
      <main className="mx-auto flex w-full max-w-7xl flex-1 flex-col px-4 py-4 md:py-8">
        {isLoading || isError ? (
          <div className="flex flex-1 items-center justify-center py-20">
            {isLoading ? (
              <CoffeeLoadingAnimation
                fullScreen={false}
                title="Admin Dashboard"
                messages={[
                  'Loading dashboard',
                  'Fetching analytics',
                  'Preparing admin tools',
                ]}
              />
            ) : (
              <CoffeeErrorAnimation
                fullScreen={false}
                title="Error Loading Dashboard"
                messages={[
                  'Failed to load dashboard',
                  'Unable to fetch data',
                  error?.message || 'An unexpected error occurred',
                ]}
              />
            )}
          </div>
        ) : (
          <Outlet />
        )}
      </main>
    </div>
  );
};

export default AdminLayout;
