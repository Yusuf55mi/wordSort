Soru Seti
Tebrikler,
İşe alım için ilk adımı geçtiniz ve teknik değerlendirme için uygun bulundunuz.
Teknik değerlendirmenin pozitif sonuçlanmasının hemen ardından işe alım süreci
başlatılacaktır.
Bu değerlendirme sürecinde danışmanınız Seray Uzgur ( seray.uzgur@fill-labs.com ) olacaktır.
Bütün sorularınızı kendisine e-posta üzerinden sorabilir, süreciniz sonlanana kadar iletişim
halinde kalabilirsiniz. Biten soruların teslimi yine danışmanınıza yapacağınız bilgilendirme ile
gerçekleşecektir.
Bitiş e-postanızı lütfen belirtilen şekilde oluşturun;
Konu 2024 Ocak - Yazılım Staj
İçerik Proje geliştirme sürecinde yaşadıklarınızı anlatan kısa bir bilgilendirme yazısı
Ek Proje klasörünün Ad-Soyad.zip olarak sıkıştırılmış hali.
Başarılı bir süreç geçirmeniz adına önemli bilgileri paylaştık, lütfen dikkatle okuyunuz.
• Yapılan çalışmanın alternatif bir durum belirtilmediği sürece 1 hafta içerisinde bitirilmesi
öngörülmektedir. Lütfen size engel olacağını düşündüğünüz bir durum varsa ya da ek süreye
ihtiyacınız varsa danışmanınız ile konuşup ek süre isteyin.
• Yapılan çalışmada 1,2 ve 3. Sorular Go dili kullanarak yazılmalıdır. 4 numaralı soru için soru
içerisinde kullanılacak olan araçlar belirtilmiştir.
• Yapılan çalışmayı Ad-Soyad.zip olarak tüm projeleri içerisinde barındıran tek bir klasör olarak
paylaşmanızı rica ediyoruz. Proje klasörleri içerisinde .git klasörünün olduğundan lütfen emin
olunuz.
• Danışmanınıza soru sormaktan ya da internet üzerinden araştırma yapmaktan çekinmeyin.
Firmamızda iyi bir araştırmacı olmak bilmek kadar değerlidir, öğrenebilmek bir beceridir.
• Herhangi bir belirsizlik bırakmamak adına değerlendirilecek kriterlerin listesi aşağıdaki gibidir:
▪ Proje geliştirme süreci
▪ GIT kullanımı
▪ Proje yapısı
▪ Değişken / Fonksiyon isimlendirmeleri (İngilizce)
▪ Test yazılması
▪ Yorum yazılması (İngilizce)
▪ Dokümantasyon (README v.b.) (İngilizce)
▪ Yeni teknolojilere uyum süreci
▪ Teslim süresi
Yakın zamanda tekrar sizden haber almak dileğiyle.

---

This part of the documentation is in English on purpose. It will help us to have an opinion about
your level of English.
Q1) Write a function that sorts a bunch of words by the number of character “a”s within the word
(decreasing order). If some words contain the same amount of character “a”s then you need to
sort those words by their lengths.
Input :
["aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"]
Output :
["aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"]
Q2) Write a recursive function which takes one integer parameter. Please bear in mind
that finding the algorithm needed to generate the output below is the main point of the
question. Please do not ask which algorithm to use.
Input :
9
Output :
2
4
9
Q3) Write a function which takes one parameter as an array/list. Find most repeated data within
a given array.
Test with different datasets.
Input :
["apple","pie","apple","red","red","red"]
Output :
"red"

---

Q4) Write a user management project which will include;
• A master view which will list all users in a data grid. This screen will assist
users with all CRUD operations. User will be able to press 3 buttons
(New,Edit,Delete). Edit and Delete operations will require row selection from the data grid.
• A detailed view which will show the fields as form. Form will have 2 buttons (Action,
Back). Text of the action button will change according to the
operation opened the detail view. For example if the “New” operation is selected from
the master, the detailed view action button text will be “Create”. Please see the
mappings below.
▪ New: Create
▪ Edit: Save
▪ Delete: Delete
• A REST service to support functions below. Please note that API paths and HTTP methods
and HTTP Statuses are important for us. Please follow REST API standards.
▪ Returns all users
▪ Return the user with the desired “id”
▪ Save the given user.
▪ Update data of the user with the desired “id”
▪ Delete the user with the desired “id”
• Backend must be written with Go. Please use sqlite for the database and include the file in
the project folder. Remember all operations must be persistent.
• Frontend must be written with TS using React & Nextjs.
