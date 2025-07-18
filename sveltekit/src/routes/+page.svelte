<script lang="js">
  import {
    searchTerm,
    searchResult,
    errorMessage,
    isLoading,
    progressPercent,
    progressMessage,
    performSearch,
    handleKeyPress,
    renderMarkdown
  } from '../lib/index';
</script>

<div class="flex flex-col items-center justify-center min-h-screen p-6 bg-gray-900 text-gray-100 font-sans">
<h1 class="text-5xl font-extrabold text-white mb-6">DeepSearch</h1>
<p class="text-lg text-gray-300 mb-8">Gerçek zamanlı arama yapmak için bir sorgu girin.</p>


<div class="flex gap-4 mb-4 w-full max-w-xl">
  <input
    type="text"
    placeholder="ewriq etc."
    bind:value={$searchTerm} 
    on:keypress={handleKeyPress}
    disabled={$isLoading}
    class="flex-grow p-3 rounded-lg border-2 border-gray-600 bg-gray-700 text-gray-100 text-lg
           focus:border-blue-500 focus:ring-2 focus:ring-blue-500 outline-none transition-all duration-300"
  />
  <button
    on:click={performSearch}
    disabled={$isLoading}
    class="bg-purple-700 text-white font-bold py-3 px-6 rounded-lg shadow-lg
           hover:bg-purple-600 active:bg-purple-800 focus:outline-none focus:ring-2 focus:ring-purple-500
           transition-all duration-300 disabled:bg-gray-600 disabled:cursor-not-allowed flex items-center justify-center gap-2"
  >
    {#if $isLoading}
      <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <span>Aranıyor...</span>
    {:else}
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8"></circle>
        <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
      </svg>
      <span>Ara</span>
    {/if}
  </button>
</div>

{#if $isLoading}
  <div class="w-full max-w-xl mt-4 mb-6 transition-opacity duration-300">

    <div class="w-full bg-gray-700 rounded-full h-2.5">
      <div 
        class="bg-purple-600 h-2.5 rounded-full transition-all duration-500 ease-out" 
        style="width: {$progressPercent}%"
      ></div>
    </div>
    <p class="text-center text-sm text-gray-400 mt-2">{$progressMessage || '...'}</p>
  </div>
{/if}

{#if $errorMessage}
  <div class="bg-red-800/50 border border-red-700 text-white font-semibold p-4 rounded-lg shadow-md mt-6 w-full max-w-xl text-center">
    {$errorMessage}
  </div>
{:else if $searchResult}
  <div class="bg-gray-800 border border-gray-700 rounded-xl p-8 mt-6 w-full max-w-3xl shadow-xl animate-fade-in">
    <h2 class="text-3xl font-semibold text-white mb-6 border-b border-gray-700 pb-4">Arama Sonucu</h2>
    <div class="rendered-markdown bg-gray-700/50 p-6 rounded-lg overflow-x-auto text-gray-200 leading-relaxed text-left">
      {@html renderMarkdown($searchResult)}
    </div>
  </div>
{/if}

</div>

<style>
/* Gerekli CSS Stilleri - Zaten mevcuttu */
.rendered-markdown p { margin-bottom: 1rem; }
.rendered-markdown h1, .rendered-markdown h2, .rendered-markdown h3 {
  margin-top: 1.5em; margin-bottom: 0.8em; font-weight: 600; border-bottom: 1px solid #4a5568; padding-bottom: 0.3em;
}
.rendered-markdown ul, .rendered-markdown ol { list-style-position: inside; padding-left: 1.5em; margin-bottom: 1em; }
.rendered-markdown li { margin-bottom: 0.5em; }
.rendered-markdown pre { background-color: #1a202c; color: #cbd5e0; padding: 1em; border-radius: 0.5rem; overflow-x: auto; margin-bottom: 1em; }
.rendered-markdown code { background-color: #2d3748; color: #f7fafc; padding: 0.2em 0.4em; border-radius: 0.25rem; }
.rendered-markdown a { color: #63b3ed; text-decoration: underline; }

/* Sonuçların yumuşak bir şekilde görünmesi için animasyon */
@keyframes fade-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in {
  animation: fade-in 0.5s ease-out forwards;
}
</style>