import { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { Link } from 'react-router-dom';
import Image from '../organism/Image';
import { usePublicCategories } from '@/features/menu/hooks/category.hook';
import { usePublicMenus } from '@/features/menu/hooks/menu.hook';
import { resolveImageUrl } from '@/utils/imageBuilder';
import type { PublicCategory } from '@/features/menu/types/category';
import type { PublicMenu } from '@/features/menu/types/menu';

import CoffeeLoadingAnimation from '@/components/shared/CoffeeLoadingAnimation';

const MenuFavoriteSection = () => {
  const [selectedCategory, setSelectedCategory] = useState('');
  const { data: categories = [] as PublicCategory[] } = usePublicCategories();

  const { data: menuItems = [] as PublicMenu[], isLoading } = usePublicMenus(
    undefined,
    selectedCategory
  );

  // Set default category to the first available one
  useEffect(() => {
    if (categories.length > 0 && !selectedCategory) {
      setSelectedCategory(categories[0].id);
    }
  }, [categories, selectedCategory]);

  const displayedMenu = menuItems.slice(0, 3);

  return (
    <section id="menu" className="bg-[#f8f3e9] py-16">
      <div className="mx-auto max-w-7xl px-4 md:px-6">
        {/* Judul dan deskripsi di tengah */}
        <div className="flex flex-col items-center justify-center space-y-4 text-center">
          <div className="space-y-2">
            <h2 className="text-3xl font-bold tracking-tighter text-[#6f4e37] sm:text-4xl md:text-5xl">
              Menu Favorit
            </h2>
            <p className="mx-auto max-w-[700px] text-[#8c7158] md:text-xl">
              Pilihan kopi dan makanan terbaik kami yang paling disukai
              pelanggan
            </p>
          </div>
        </div>

        {/* Kategori menu */}
        <div className="mt-8 mb-12 grid grid-cols-2 gap-2 sm:grid-cols-3 sm:gap-3 md:grid-cols-4 md:gap-4 lg:flex lg:flex-wrap lg:justify-center lg:gap-4">
          {categories.map(category => (
            <Button
              key={category.id}
              variant={selectedCategory === category.id ? 'default' : 'outline'}
              className={`w-full cursor-pointer rounded-full px-3 py-2 text-center text-xs transition-all sm:px-4 sm:py-2 sm:text-sm md:px-5 md:py-2 md:text-sm lg:w-auto lg:px-6 lg:py-2 lg:text-base ${
                selectedCategory === category.id
                  ? 'border border-transparent bg-[#6f4e37] text-white hover:bg-[#5d4130]'
                  : 'border border-[#6f4e37] text-[#6f4e37] hover:bg-[#6f4e37] hover:text-white'
              } `}
              onClick={() => setSelectedCategory(category.id)}
            >
              {category.name}
            </Button>
          ))}
        </div>

        {/* Kartu menu */}
        {isLoading ? (
          <CoffeeLoadingAnimation
            fullScreen={false}
            title="Mencari Menu"
            messages={['Melihat stok', 'Mempersiapkan rasa', 'Tunggu sebentar']}
          />
        ) : (
          <div className="mt-12 grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
            {displayedMenu.map(item => (
              <div
                key={item.id}
                className="group relative overflow-hidden rounded-lg border border-[#e6d9c9] bg-white shadow-sm transition-all hover:shadow-md"
              >
                <div className="aspect-square overflow-hidden">
                  <Image
                    image_url={resolveImageUrl(item.product_image_file, {
                      w: 400,
                      h: 400,
                      q: 80,
                    })}
                    alt_text={item.name}
                    width={400}
                    height={400}
                    class_name="h-full w-full object-cover transition-transform group-hover:scale-105"
                  />
                </div>
                <div className="p-6">
                  <div className="mb-2 flex items-center justify-between">
                    <span className="rounded-full bg-[#f8f3e9] px-2 py-1 text-xs font-medium text-[#a67c52]">
                      {item.category.name}
                    </span>
                  </div>
                  <h3 className="text-xl font-bold text-[#6f4e37]">
                    {item.name}
                  </h3>
                  <div className="mt-4 flex items-center justify-between">
                    <span className="text-lg font-bold text-[#6f4e37]">
                      Rp {Intl.NumberFormat('id-ID').format(item.price)}
                    </span>
                    <span className="rounded-full bg-[#f8f3e9] px-3 py-1 text-sm text-[#8c7158]">
                      {item.stock > 0 ? 'Tersedia' : 'Habis'}
                    </span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}

        {/* Tombol aksi */}
        <div className="mt-12 text-center">
          <Link to="/menu-lengkap">
            <Button className="cursor-pointer bg-[#6f4e37] text-white hover:bg-[#5d4130]">
              Lihat Menu Lengkap
            </Button>
          </Link>
        </div>
      </div>
    </section>
  );
};

export default MenuFavoriteSection;
