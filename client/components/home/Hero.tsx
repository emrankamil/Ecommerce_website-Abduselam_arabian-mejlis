import initTranslations from "@/app/i18n";

export default async function Hero({ params: { locale } }: { params: { locale: string } }) {
  const {t} = await initTranslations(locale, ['home'])
  return (
    <div className="relative bg-[url('/hero-image.png')] bg-cover bg-left h-screen">
      <div className="absolute inset-0 bg-black bg-opacity-40 rounded-xl"></div>
      <div className="flex h-16 items-center justify-between px-4 sm:px-6 lg:px-8"></div>
    </div>
  );
}

