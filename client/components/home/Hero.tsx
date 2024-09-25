import initTranslations from "@/app/i18n";
import { Button } from "../ui/button";
import Image from "next/image";
import Link from "next/link";

export default async function Hero({
  params: { locale },
}: {
  params: { locale: string };
}) {
  const { t } = await initTranslations(locale, ["home"]);
  return (
    <div className="relative bg-[url('/hero-image.png')] bg-cover bg-left h-screen">
      <div className="absolute bottom-0 left-1/2 z-10">
        <Link href="#produts" passHref>
          <Image src="/drag.svg" alt={""} width={25} height={25} />
        </Link>
      </div>
      <div className="absolute inset-0 bg-gradient-to-tr from-black/100 to-transparent bg-opacity-70 rounded-xl" />
      <div className="relative w-1/2 h-full  z-10 flex flex-col justify-center px-24 py-6 w-2/3 space-y-8">
        <h1 className="text-5xl text-white leading-normal">
          Abduselam Arabian Mejlis Manufacturing Experts
        </h1>
        <p className="text-lg text-white">
          We are not just suppliers, we are your partners <br />
          Abduselam members
        </p>
        <div className="flex space-x-20">
          <Button
            variant="outline"
            className="mt-6 px-8 py-7 w-48 text-primary rounded-full font-semibold border-2 border-white/50 bg-white/90 hover:bg-gray-200/20 transition duration-300"
          >
            SHOP NOW
          </Button>

          <Button
            variant="outline"
            className="mt-6 px-8 py-7 w-48 text-white rounded-full font-semibold border-2 border-white/50 bg-transparent hover:bg-gray-200/20 transition duration-300"
          >
            LET'S TALK
          </Button>
        </div>
      </div>
    </div>
  );
}
