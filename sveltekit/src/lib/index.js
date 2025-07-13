// place files you want to import through the `$lib` alias in this folder.
// ./lib/search.js
import { writable } from 'svelte/store';
import axios from 'axios';
import { marked } from 'marked';

// Store tanımları
export const searchTerm = writable('');
export const searchResult = writable(null);
export const errorMessage = writable(null);
export const isLoading = writable(false);

// Arama fonksiyonu
export async function performSearch() {
  errorMessage.set(null);
  searchResult.set(null);
  isLoading.set(true);

  let term = '';
  const unsubscribe = searchTerm.subscribe(value => {
    term = value;
  });
  unsubscribe(); // Aboneliği hemen iptal et, çünkü sadece o anki değeri istiyoruz

  if (term.trim() === '') {
    errorMessage.set('Lütfen bir arama terimi girin.');
    isLoading.set(false);
    return;
  }

  try {
    // DÜZELTME: API URL'sine '/api' eklendi
    const response = await axios.get(`http://localhost:3000/search/${encodeURIComponent(term)}`);
    const data = response.data;

    if (data.status === 'success') {
      searchResult.set(data.data || 'Sonuç bulunamadı.');
    } else if (data.status === 'no content') {
      errorMessage.set(data.message || 'Arama terimi eksik.');
    } else { // status === 'error'
      errorMessage.set(data.message || 'Arama sırasında bir hata oluştu.');
    }
  } catch (error) {
    if (axios.isAxiosError && axios.isAxiosError(error)) { // axios.isAxiosError kontrolü zaten bir bool döndürür, çift kontrol gereksiz
      if (error.response) {
        errorMessage.set(error.response.data?.message || `API Hatası: ${error.response.status}`);
      } else if (error.request) {
        errorMessage.set('Sunucuya ulaşılamıyor. Ağ bağlantınızı kontrol edin.');
      } else {
        errorMessage.set(error.message || 'Bilinmeyen bir Axios hatası oluştu.');
      }
    } else if (error instanceof Error) {
      errorMessage.set(error.message);
    } else {
      errorMessage.set('Bilinmeyen bir hata oluştu.');
    }
    console.error('API hatası:', error);
  } finally {
    isLoading.set(false);
  }
}

// Enter tuşu ile aramayı tetikleyen fonksiyon
export function handleKeyPress(event) {
  if (event.key === 'Enter') {
    performSearch();
  }
}

// Markdown render fonksiyonu
export function renderMarkdown(text) {
    if (text === null || text === undefined) { // Null veya undefined kontrolü ekle
        return '';
    }
  return marked(text);
}