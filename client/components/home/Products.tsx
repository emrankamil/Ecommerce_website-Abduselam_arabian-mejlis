"use client";

import { useState } from "react";
import Image from "next/image";
import { FaStar } from "react-icons/fa6";
import { Button } from "../ui/button";

const products: { [key: string]: string[] } = {
  Popular: [
    "/products/popular-1.png",
    "/products/popular-2.png",
    "/products/popular-3.png",
    "/products/popular-2.png",
  ],
  "Arabian Majlis": [
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
  "Tv Stand": [
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
  const [fadeIn, setFadeIn] = useState(true);

  const handleCategoryChange = (category: string) => {
    setFadeIn(false); // Trigger fade-out animation
    setTimeout(() => {
      setSelectedCategory(category);
      setImages(products[category]);
      setFadeIn(true); // Trigger fade-in animation
    }, 300);
  };

  return (
    <div id="products" className="mx-auto px-10">
      <h3 className="text-xl font-bold text-center mb-6">Products</h3>
      <h2 className="text-4xl text-center mb-6">Magnificent Product from Us</h2>

      {/* Category Tabs */}
      <div className="flex justify-center space-x-10 mb-10">
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
      <div
        className={`grid grid-cols-2 pb-6 transition-opacity duration-500 ease-in-out ${
          fadeIn ? "opacity-100" : "opacity-0"
        }`}
      >
        {images.map((image, index) => (
          <div
            key={index}
            className="overflow-hidden transform transition-transform duration-500 p-6 h-[75vh]"
            style={{ transition: "transform 0.5s ease-in-out" }}
          >
            {index === 0 ? (
              <div className="max-w-4xl mx-auto bg-white shadow-lg h-full">
                {/* Image Section */}
                <div className="overflow-hidden h-2/3">
                  <img
                    src={image}
                    alt="Product"
                    className="w-full h-full object-cover"
                  />
                </div>

                {/* Info Section */}
                <div className="p-4 flex justify-between h-1/3">
                  <div className="w-2/3">
                    {/* Title */}
                    <h2 className="text-2xl font-bold mb-2">Arabian Mejlis</h2>
                    {/* Description */}
                    <p className="text-gray-700 mb-4">
                      Indulge in comfort and style with our premium sofas,
                      designed to be the centerpiece of your living space.
                    </p>
                  </div>

                  <div className="flex-col justify-between  w-1/3">
                    <div className="text-center">
                      {/* Rating */}
                      <h3 className="text-lg font-bold ">Best Selling</h3>
                      <div className="flex items-center justify-center space-x-1">
                        {[...Array(5)].map((_, i) => (
                          <FaStar
                            key={i}
                            className="text-yellow-500 text-lg mb-2"
                          />
                        ))}
                      </div>
                    </div>

                    {/* Show Details Button */}
                    <div className="flex items-center justify-center">
                      <Button className="bg-black text-white px-4 py-2 rounded-xl hover:bg-gray-800 transition duration-300">
                        SHOW DETAILS
                      </Button>
                    </div>
                  </div>
                </div>
              </div>
            ) : (
              <Image
                src={image}
                alt={`Category ${selectedCategory} image ${index + 1}`}
                width={500}
                height={1000}
                className="w-full h-full object-cover "
              />
            )}
          </div>
        ))}
      </div>
      <div className="w-full flex justify-center ">
        <Button
          variant={"outline"}
          className="font-semibold px-20 py-6 rounded-xl hover:bg-gray-200 transition duration-300"
        >
          SEE MORE
        </Button>
      </div>
    </div>
  );
};

export default ProductShowcase;
