import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'motion/react';
import { Coffee } from 'lucide-react';

interface CoffeeLoadingAnimationProps {
  messages?: string[];
  interval?: number;
  title?: string;
  fullScreen?: boolean;
}

const CoffeeLoadingAnimation = ({
  messages = ['Menyeduh kopi...', 'Mempersiapkan aroma...', 'Hampir siap...'],
  interval = 3000,
  title = 'Sedang Memproses',
  fullScreen = true,
}: CoffeeLoadingAnimationProps) => {
  const [index, setIndex] = useState(0);

  useEffect(() => {
    if (messages.length > 0) {
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
        {/* Minimalist Pulsing Icon */}
        <motion.div
          animate={{ opacity: [0.4, 1, 0.4] }}
          transition={{ duration: 2, repeat: Infinity, ease: 'easeInOut' }}
          className="mb-8"
        >
          <Coffee className="h-10 w-10 text-[#6f4e37]" />
        </motion.div>

        {/* Clean Typography */}
        <div className="text-center">
          <h2 className="mb-2 text-[10px] font-bold tracking-[0.4em] text-[#6f4e37] uppercase">
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

        {/* Ultra-thin Minimalist Progress Line */}
        <div className="mt-8 h-[1px] w-32 overflow-hidden bg-[#e6d9c9]">
          <motion.div
            initial={{ x: '-100%' }}
            animate={{ x: '100%' }}
            transition={{ duration: 2.5, repeat: Infinity, ease: 'linear' }}
            className="h-full w-full bg-[#6f4e37]"
          />
        </div>
      </div>
    </div>
  );
};

export default CoffeeLoadingAnimation;
