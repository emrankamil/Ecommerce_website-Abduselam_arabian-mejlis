"use client";

import { useState } from "react";
import Image from "next/image";

const products: { [key: string]: string[] } = {
  Popular: [
    "/products/popular-1.png",
    "/products/popular-2.png",
    "/products/popular-3.png",
    "/products/popular-2.png",
  ],
  Arabian_Majlis: [
    "/products/popular-1.png",
    "/products/popular-2.png",
    "/products/popular-3.png",
    "/products/popular-4.png",
  ],
  Sofa: [
    "/products/popular-1.png",
    "/products/popular-2.png",
    "/products/popular-3.png",
    "/products/popular-4.png",
  ],
  Curtains: [
    "/products/popular-1.png",
    "/products/popular-2.png",
    "/products/popular-3.png",
    "/products/popular-4.png",
  ],
  Beds: [
    "/products/popular-1.png",
    "/products/popular-2.png",
    "/products/popular-3.png",
    "/products/popular-4.png",
  ],
  Tv_Stand: [
    "/products/popular-1.png",
    "/products/popular-2.png",
    "/products/popular-3.png",
    "/products/popular-4.png",
  ],
};

const ProductShowcase = () => {
  const [selectedCategory, setSelectedCategory] =
    useState<keyof typeof products>("Popular");
  const [images, setImages] = useState<string[]>(products[selectedCategory]);

  const handleCategoryChange = (category: string) => {
    setSelectedCategory(category);
    setImages(products[category]);
  };

  return (
    <div className="container mx-auto py-10">
      {/* Title */}
      <h2 className="text-4xl font-bold text-center mb-6">
        Magnificent Product from Us
      </h2>

      {/* Category Tabs */}
      <div className="flex justify-center space-x-6 mb-10">
        {Object.keys(products).map((category) => (
          <button
            key={category}
            className={`text-lg font-medium ${
              selectedCategory === category
                ? "text-black border-b-2 border-black"
                : "text-gray-400"
            } transition-all duration-300 ease-in-out`}
            onClick={() => handleCategoryChange(category)}
          >
            {category}
          </button>
        ))}
      </div>

      {/* Images with transition */}
      <div className="grid grid-cols-2 gap-4">
        {images.map((image, index) => (
          <div
            key={index}
            className="overflow-hidden rounded-lg shadow-lg"
            style={{ transition: "transform 0.5s ease-in-out" }}
          >
            <Image
              src={image}
              alt={`Category ${selectedCategory} image ${index + 1}`}
              width={500}
              height={300}
              className="transform hover:scale-105 transition-transform duration-500"
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default ProductShowcase;
