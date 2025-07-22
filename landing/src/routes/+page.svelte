<script>
    import { Github } from "lucide-svelte";

    function clicked() {
        window.location.href = "https://github.com/ewriq/deepsearch";
    }
</script>

<div
    class="bg-gradient-to-b from-gray-900 to-purple-900 text-white min-h-screen font-sans"
>
    <header
        class="container mx-auto py-4 px-6 flex justify-between items-center"
    >
        <span class="font-semibold text-xl">DeepSearch</span>
        <div class="flex items-center space-x-4">
            <button
                on:click={clicked}
                class="bg-purple-700 hover:bg-purple-600 text-white font-semibold py-2 px-4 rounded-md text-sm"
            >
                <Github class="h-5 w-5" />
            </button>
        </div>
    </header>

    <main
        class="container mx-auto flex flex-col items-center py-16 px-6 space-y-24"
    >
        <!-- BÖLÜM 1: Analyze -->
        <section
            class="grid grid-cols-1 md:grid-cols-2 gap-8 items-center w-full mt-40"
        >
            <div>
                <h2 class="text-2xl font-bold mb-4">
                    Analyze in Your <span class="text-red-500">QUERY'S</span>
                </h2>
                <p class="text-gray-400 mb-6 text-sm">
                    Lorem Ipsum<br />
                    "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet,
                    consectetur, adipisci velit..."<br />
                    "There is no one who loves pain itself, who seeks after it and
                    wants to have it, simply because it is pain..."
                </p>
                <button
                    on:click={clicked}
                    class="bg-purple-700 hover:bg-purple-600 text-white font-semibold py-2 px-4 rounded-md text-sm"
                >
                    Quick Explore
                </button>
            </div>
            <div class="flex justify-end">
                <img
                    src="/thinkink.png"
                    alt="Thinking Illustration"
                    class="rounded-lg max-w-full h-auto"
                />
            </div>
        </section>

        <div class="h-[500px] mt-160">
            <section
                class="grid grid-cols-1 md:grid-cols-2 gap-8 items-center w-full"
            >
                <div class="flex justify-end md:order-last">
                    <img
                        src="/favicon.png"
                        alt="Water Drop Illustration"
                        class="rounded-full max-w-full h-auto w-[250px] md:w-[300px]"
                    />
                </div>
                <div class="md:order-first">
                    <h2 class="text-2xl font-bold mb-4">
                        Research Your <span class="text-blue-500"
                            >Knowledge</span
                        >
                    </h2>
                    <p class="text-gray-400 mb-6 text-sm">
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit,
                        sed do eiusmod tempor incididunt ut labore et dolore
                        magna aliqua. Ut enim ad minim veniam, quis nostrud
                        exercitation ullamco laboris nisi ut aliquip ex ea
                        commodo consequat.
                    </p>
                    <button
                        on:click={clicked}
                        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-md text-sm"
                    >
                        Learn More
                    </button>
                </div>
            </section>
        </div>
    </main>

    <!-- DÜZELTME: Bu bölüme mt-24 eklenerek ana içerikten belirgin şekilde ayrıldı. -->
    <div class="flex flex-col items-center mt-24 mb-4 px-6 mt-60">
        <Github class="h-8 w-8 mb-2" />
        <h3 class="text-xl font-semibold">
            GitHub <a
                href="https://github.com/ewriq/deepsearch"
                class="font-bold text-white hover:underline">DeepSearch</a
            >
        </h3>
        <p class="text-sm text-center">
            Search For The Ones You Have In Mind And You Will Have The Final
            Result
        </p>
    </div>

    <div class="flex items-center justify-center space-x-2 mb-8">
        <img
            class="w-[50px] h-[50px] rounded-full"
            src="https://avatars.githubusercontent.com/u/102306313?v=4"
            alt="User Avatar"
        />
        <a
            href="https://github.com/ewriq/"
            class="text-sm text-red-500 hover:underline">@ewriq</a
        >
    </div>

    <!-- DÜZELTME: Footer'a mt-16 eklenerek üstündeki bölümle arasına mesafe konuldu. -->
    <footer
        class="container mx-auto py-10 px-6 text-center text-gray-400 mt-16"
    >
        <p class="text-xs">
            Copyright 2025 for <span class="text-red-500">@ewriq</span>
        </p>
    </footer>
</div>
