import { lucideIcons } from '@/icon/lucide-react-icons';
const { Search } = lucideIcons;
import type { PublicMenu } from '../../types/menu';
import { resolveImageUrl } from '@/utils/imageBuilder';

interface MenuGridProps {
  filteredMenuItems: PublicMenu[];
}

const MenuGrid = ({ filteredMenuItems }: MenuGridProps) => {
  return (
    <section className="bg-[#e6d9c9] py-12">
      <div className="mx-auto max-w-7xl px-6">
        {filteredMenuItems.length === 0 ? (
          <div className="py-16 text-center">
            <div className="mx-auto mb-6 flex h-24 w-24 items-center justify-center rounded-full bg-white/70 shadow-lg">
              <Search className="h-12 w-12 text-[#6f4e37]" />
            </div>
            <h3 className="mb-3 text-2xl font-bold text-[#6f4e37]">
              Menu tidak ditemukan
            </h3>
            <p className="mx-auto max-w-md text-lg text-[#8c7158]">
              Coba kata kunci lain untuk menemukan menu favorit Anda
            </p>
          </div>
        ) : (
          <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
            {filteredMenuItems.map(item => (
              <div
                key={item.id}
                className="group relative overflow-hidden rounded-2xl bg-white shadow-md transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
              >
                {/* Image Container */}
                <div className="relative aspect-[4/3] overflow-hidden rounded-t-2xl">
                  <img
                    src={resolveImageUrl(item.product_image_file, {
                      w: 400,
                      h: 300,
                      q: 80,
                    })}
                    alt={item.name}
                    className="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
                  />
                </div>

                {/* Content */}
                <div className="space-y-3 p-4">
                  {/* Category */}
                  <span className="inline-block rounded-full bg-[#f0e8dc] px-3 py-1 text-sm font-medium text-[#6f4e37]">
                    {item.category.name}
                  </span>

                  {/* Menu Name */}
                  <h3 className="line-clamp-2 text-lg leading-snug font-bold text-[#6f4e37]">
                    {item.name}
                  </h3>

                  {/* Price */}
                  <div className="flex items-center justify-between">
                    <span className="text-2xl font-bold text-[#6f4e37]">
                      Rp {Intl.NumberFormat('id-ID').format(item.price)}
                    </span>
                  </div>

                  {/* Stock Status */}
                  <div className="flex items-center justify-between pt-2">
                    <span className="text-sm text-[#8c7158]">
                      Stok: {item.stock}
                    </span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </section>
  );
};

export default MenuGrid;
