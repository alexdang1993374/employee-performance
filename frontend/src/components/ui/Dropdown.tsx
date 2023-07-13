import { AnimatePresence, motion } from "framer-motion";
import React, { useEffect, useRef, useState } from "react";
import { useMediaQuery } from "react-responsive";

import { MenuItem } from "@/types";

interface DropdownProps {
  buttonLabel: string;
  menuItems: MenuItem[];
}

const Dropdown: React.FC<DropdownProps> = ({ buttonLabel, menuItems }) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const wrapperRef = useRef<HTMLDivElement>(null);
  const isMobile = useMediaQuery({ maxWidth: 767 });

  const handleClick = () => {
    if (isMobile) {
      setIsOpen(!isOpen);
    }
  };

  const handleMouseEnter = () => {
    if (!isMobile) {
      setIsOpen(true);
    }
  };

  const handleMouseLeave = () => {
    if (!isMobile) {
      setIsOpen(false);
    }
  };

  const handleItemClick = (onClick: VoidFunction) => {
    onClick();
    setIsOpen(false);
  };

  useEffect(() => {
    const handleOutsideClick = (event: MouseEvent) => {
      if (
        isMobile &&
        isOpen &&
        wrapperRef.current &&
        !wrapperRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleOutsideClick);
    return () => {
      document.removeEventListener("mousedown", handleOutsideClick);
    };
  }, [isMobile, isOpen]);

  return (
    <div
      className="relative inline-block"
      ref={wrapperRef}
      onMouseEnter={handleMouseEnter}
      onMouseLeave={handleMouseLeave}
    >
      <button
        className={`w-[150px] flex justify-center items-center bg-gray-500 text-white px-3 py-2 rounded-full text-sm ${
          isOpen ? "opacity-70" : "opacity-100"
        }`}
        onClick={handleClick}
      >
        {buttonLabel}
      </button>

      <AnimatePresence>
        {isOpen && (
          <motion.div
            className="absolute bg-gray-200 rounded text-center shadow-xl overflow-hidden z-10 w-[150px]"
            initial={{ height: 0 }}
            animate={{ height: "auto" }}
            exit={{ height: 0 }}
            transition={{ duration: 0.2, ease: "linear" }}
          >
            {menuItems.map((item, index) => (
              <div
                key={index}
                className="text-black py-3 px-4 cursor-pointer hover:bg-gray-300 text-sm"
                onClick={() => handleItemClick(item.onClick)}
              >
                {item.label}
              </div>
            ))}
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
};

export default Dropdown;
