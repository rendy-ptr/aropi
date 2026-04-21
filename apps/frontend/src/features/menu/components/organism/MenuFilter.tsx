import { usePublicCategories } from '../../hooks/category.hook';
import type { PublicCategory } from '../../types/category';

interface MenuFilterProps {
  selectedCategory: string;
  setSelectedCategory: (id: string) => void;
}

const MenuFilter = ({
  selectedCategory,
  setSelectedCategory,
}: MenuFilterProps) => {
  const { data: categories = [], isLoading, error } = usePublicCategories();

  if (isLoading) {
    return (
      <div className="flex justify-center py-4">
        <div className="flex items-center gap-2 text-[#6f4e37]">
          <div className="h-5 w-5 animate-spin rounded-full border-2 border-[#6f4e37]/20 border-t-[#6f4e37]"></div>
          <span className="text-sm font-medium">Memuat kategori...</span>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex justify-center py-4">
        <div className="rounded-lg border border-red-200 bg-red-50 px-4 py-2 text-center text-red-500">
          <p className="text-sm font-medium">⚠️ Gagal memuat kategori</p>
        </div>
      </div>
    );
  }

  const categoriesWithAll: PublicCategory[] = [
    { id: 'semua', name: 'Semua Menu' },
    ...categories,
  ];

  return (
    <div className="w-full">
      {/* Category Label */}
      <div className="mb-4 text-center">
        <span className="rounded-full bg-white/50 px-3 py-1 text-sm font-semibold text-[#6f4e37]">
          Kategori Menu
        </span>
      </div>

      {/* Categories Grid */}
      <div className="mx-auto flex max-w-5xl flex-wrap justify-center gap-2 sm:gap-3">
        {categoriesWithAll.map(category => {
          const isActive = selectedCategory === category.id;

          return (
            <button
              key={category.id}
              onClick={() => setSelectedCategory(category.id)}
              className={`group relative flex transform items-center gap-2 rounded-xl px-4 py-2.5 text-sm font-semibold transition-all duration-300 ease-out hover:scale-105 active:scale-95 ${
                isActive
                  ? 'bg-[#6f4e37] text-white shadow-lg ring-2 shadow-[#6f4e37]/25 ring-[#6f4e37]/20'
                  : 'border-2 border-[#e6d9c9]/50 bg-white/80 text-[#6f4e37] backdrop-blur-sm hover:border-[#d4c3b0] hover:bg-white hover:shadow-md hover:shadow-[#6f4e37]/10'
              } min-w-0 flex-shrink-0`}
              aria-pressed={isActive}
            >
              {/* Text */}
              <span className="relative z-10 whitespace-nowrap">
                {category.name}
              </span>
            </button>
          );
        })}
      </div>

      {/* Active Filter Indicator */}
      {selectedCategory !== 'semua' && (
        <div className="mt-4 text-center">
          <div className="inline-flex items-center gap-2 rounded-full bg-[#6f4e37]/10 px-3 py-1 text-xs font-medium text-[#6f4e37]">
            <div className="h-2 w-2 animate-pulse rounded-full bg-[#6f4e37]"></div>
            Filter aktif:{' '}
            {categoriesWithAll.find(cat => cat.id === selectedCategory)?.name}
          </div>
        </div>
      )}
    </div>
  );
};

export default MenuFilter;
