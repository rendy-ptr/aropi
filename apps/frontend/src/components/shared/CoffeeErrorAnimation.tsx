import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'motion/react';
import { Coffee, RotateCcw } from 'lucide-react';

interface CoffeeErrorAnimationProps {
  messages?: string[];
  interval?: number;
  title?: string;
  onRetry?: () => void;
  fullScreen?: boolean;
}

const CoffeeErrorAnimation = ({
  messages = [
    'Oops! Gangguan teknis',
    'Koneksi terputus.',
    'Silakan coba lagi.',
  ],
  interval = 4000,
  title = 'Gagal Memuat',
  onRetry,
  fullScreen = true,
}: CoffeeErrorAnimationProps) => {
  const [index, setIndex] = useState(0);

  useEffect(() => {
    if (messages.length > 1) {
      const timer = setInterval(() => {
        setIndex(prev => (prev + 1) % messages.length);
      }, interval);
      return () => clearInterval(timer);
    }
  }, [messages.length, interval]);

  return (
    <div
      className={`flex items-center justify-center ${
        fullScreen ? 'min-h-screen bg-[#f8f3e9]' : 'py-20'
      }`}
    >
      <div className="flex flex-col items-center">
        {/* Minimalist Tilted Error Icon */}
        <motion.div
          initial={{ rotate: 0 }}
          animate={{ rotate: -15 }}
          className="mb-8"
        >
          <Coffee className="h-10 w-10 text-red-800 opacity-60" />
        </motion.div>

        {/* Clean Typography */}
        <div className="text-center">
          <h2 className="mb-2 text-[10px] font-bold tracking-[0.4em] text-red-900 uppercase">
            {title}
          </h2>

          <div className="h-6">
            <AnimatePresence mode="wait">
              <motion.p
                key={index}
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
                transition={{ duration: 0.8 }}
                className="text-sm font-medium tracking-wide text-[#8c7158]"
              >
                {messages[index]}
              </motion.p>
            </AnimatePresence>
          </div>
        </div>

        {/* Minimalist Action Button */}
        <div className="mt-10">
          {onRetry && (
            <motion.button
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              onClick={onRetry}
              className="flex items-center space-x-2 border-b border-[#6f4e37]/30 pb-0.5 text-xs font-bold tracking-widest text-[#6f4e37] uppercase transition-colors hover:border-[#6f4e37]"
            >
              <RotateCcw className="h-3 w-3" />
              <span>Coba Lagi</span>
            </motion.button>
          )}
        </div>
      </div>
    </div>
  );
};

export default CoffeeErrorAnimation;
