# 👻 DeepSearch

DeepSearch, kullanıcıların çeşitli kaynaklardan bilgi araması yapmasını sağlayan, LLM (Large Language Model) destekli bir arama motoru uygulamasıdır. Bu proje, **GoFiber** framework'ü ile geliştirilmiş bir backend, **Svelte** tabanlı bir frontend ve **PostgreSQL** veritabanı kullanılarak oluşturulmuştur. Ayrıca, Tailwind CSS gibi modern araçlarla zengin bir kullanıcı arayüzü sunar.

---

## 🚀 Özellikler

- **LLM Destekli Özetleme ve Analiz**: Arama sonuçlarını özetler ve analiz eder.
- **Çoklu Arama Motoru Desteği**: Google, Yandex ve Bing gibi arama motorlarından veri çekme.
- **PostgreSQL Veritabanı**: Arama sonuçlarını depolamak ve yönetmek için kullanılır.
- **Fiber Framework**: Hızlı ve ölçeklenebilir bir backend.
- **Svelte Frontend**: Kullanıcı dostu ve performanslı bir arayüz.
- **Tailwind CSS**: Modern ve şık bir tasarım.



## 🛠️ Kurulum

### Gereksinimler

- **Go** (v1.19+)
- **Node.js** (v16+)
- **PostgreSQL** (v13+)

### Adımlar

1. **Depoyu Klonlayın**:
   ```bash
   git clone https://github.com/kullanici-adi/deepsearch.git
   cd deepsearch
   ```

2. **Backend Bağımlılıklarını Yükleyin**:
   ```bash
   go mod tidy
   ```

3. **Frontend Bağımlılıklarını Yükleyin**:
   ```bash
   cd web
   npm install
   ```

4. **Veritabanını Ayarlayın**:
   PostgreSQL'de bir veritabanı oluşturun ve `config/server.ini` dosyasındaki `dsn` değerini güncelleyin:
   ```ini
   [db]
   dsn="postgresql://user:password@localhost:5432/deepsearch"
   ```

5. **Arama Motoru ve LLM API Anahtarlarını Ayarlayın**:
   `config/search.ini` dosyasındaki `key` ve `gemini` alanlarını doldurun:
   ```ini
   [serpapi]
   key = "YOUR_SERPAPI_KEY"

   [ai]
   gemini = "YOUR_LLM_API_KEY"
   ```

6. **Backend'i Çalıştırın**:
   ```bash
   go run main.go
   ```

7. **Frontend'i Çalıştırın**:
   ```bash
   cd web
   npm run dev
   ```

8. **Uygulamayı Açın**:
   Tarayıcınızda `http://localhost:3000` adresine gidin.


## 🔧 Yapılandırma

### `server.ini`
API ve veritabanı ayarlarını içerir:
```ini
[api]
port = ":3000"

[db]
dsn = "postgresql://user:password@localhost:5432/deepsearch"
```

### `search.ini`
Arama motoru ve LLM ayarlarını içerir:
```ini
[serpapi]
key = "YOUR_SERPAPI_KEY"
google = true
yandex = true
bing = true

[ai]
gemini = "YOUR_GEMINI_API_KEY"
prompt = "Buradaki veriye göre bana bir özet çıkarmanı ve analiz etmeni..."
```


## 📜 Kullanım

1. **Arama Yapın**: Ana sayfada bir sorgu girin ve "Ara" butonuna tıklayın.
2. **Sonuçları Görüntüleyin**: Arama sonuçları özetlenmiş ve analiz edilmiş şekilde görüntülenir.
3. **Hata Mesajları**: Eğer bir hata oluşursa, kullanıcıya uygun bir mesaj gösterilir.


## 🤝 Katkıda Bulunun

Katkıda bulunmak isterseniz, lütfen bir **pull request** gönderin veya bir **issue** açın. Her türlü katkı memnuniyetle karşılanır!


## 🌟 Destek

Eğer bu projeyi beğendiyseniz, lütfen ⭐ vererek destek olun! 😊
