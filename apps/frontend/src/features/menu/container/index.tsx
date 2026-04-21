import { useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import { FOOTER_LINKS } from '@/constants/footerLinks';
import Footer from '@/components/shared/Footer';
import Navbar from '@/components/shared/Navbar';
import SearchFilterMenu from '../components/sections/Search & Filter Menu';
import MenuGrid from '../components/sections/MenuGrid';
import { usePublicMenus } from '../hooks/menu.hook';
import CoffeeLoadingAnimation from '@/components/shared/CoffeeLoadingAnimation';
import { useDebounce } from '@/hooks/useDebounce';

const UserMenuUI = () => {
  const [searchParams, setSearchParams] = useSearchParams();

  const selectedCategory = searchParams.get('category') || 'semua';
  const searchQuery = searchParams.get('search') || '';

  const setSelectedCategory = (category: string) => {
    setSearchParams(
      prev => {
        if (category === 'semua' || !category) {
          prev.delete('category');
        } else {
          prev.set('category', category);
        }
        return prev;
      },
      { replace: true }
    );
  };

  const setSearchQuery = (query: string) => {
    setSearchParams(
      prev => {
        if (!query) {
          prev.delete('search');
        } else {
          prev.set('search', query);
        }
        return prev;
      },
      { replace: true }
    );
  };

  const debouncedSearch = useDebounce(searchQuery, 1000);

  const {
    data: menuItems = [],
    isLoading,
    error,
  } = usePublicMenus(debouncedSearch, selectedCategory);

  useEffect(() => {
    window.scrollTo(0, 0);
  }, []);

  if (isLoading) {
    return (
      <CoffeeLoadingAnimation
        title="Loading Menu"
        messages={[
          'Mengambil data Menu',
          'Memproses informasi',
          'Mempersiapkan tampilan',
        ]}
      />
    );
  }

  if (error) {
    return (
      <CoffeeLoadingAnimation
        title="Gagal Memuat Menu"
        messages={[
          'Gagal memuat menu',
          'Terjadi kesalahan saat mengambil data',
          'Silahkan coba lagi',
        ]}
      />
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-[#faf7f2] to-[#f5f0e8]">
      <Navbar variant="subpage" />

      {/* Search & Filter dengan light variant */}
      <SearchFilterMenu
        searchQuery={searchQuery}
        setSearchQuery={setSearchQuery}
        selectedCategory={selectedCategory}
        setSelectedCategory={setSelectedCategory}
      />

      {/* Menu Grid dengan card premium */}
      <MenuGrid filteredMenuItems={menuItems} />

      {/* Footer */}
      <Footer variant="light" links={FOOTER_LINKS} />
    </div>
  );
};

export default UserMenuUI;
