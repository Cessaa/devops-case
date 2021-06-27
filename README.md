### Senaryo

devops-case, Kubernetes ortamında çalışması planlanan, Go ile geliştirilmiş örnek bir mikro servistir. Default olarak 9000 portunu dinleyen projenin önünde konumlanan ingress ile "dc/v1/" path ve parametrelerinin projeye route edilmesi istenmektedir. 

### Gereklilikler

1. Dockerfile oluşturulması,
2. Kubernetes ortamı için gerekli dosyalar,
3. Çalışma ortamı ile ilgili özet bir döküman (1 veya 2 paragraftan oluşabilir)

### Çalıştırma

Bu proje GKE kullanılacak şekilde hazırlandı.
Deploy scriptini calıştırmadan önce google cloud'a login olmak gerekiyor.

```bash
gcloud auth login
```

Eğer birden çok proje varsa birini set ediyoruz.

```bash
gcloud config set project PROJECT_ID
```

Son olarak deploy scriptini çalıştırıyoruz.

```bash
./deploy.sh CLUSTER_NAME CLUSTER_REGION_NAME
```

Örnek:
```bash
./deploy.sh case us-central1-c
```
