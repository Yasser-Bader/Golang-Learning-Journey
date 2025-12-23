# 🏗️ Interfaces (Notification System)

## 📝 وصف المهمة

بناء نظام مرن لإرسال الإشعارات للمستخدمين عبر قنوات مختلفة (Email و SMS) باستخدام Interface موحد، بحيث يمكننا إضافة قنوات جديدة مستقبلاً دون تغيير الكود الأساسي.

## ​⚙️ المتطلبات التقنية (Technical Requirements):

- ​تعريف الـ Interface:
  ​أنشئ interface اسمه Notifier.
  ​يحتوي على دالة واحدة فقط: Send(message string).

- ​تنفيذ القنوات (Structs):
  ​أنشئ Struct اسمه EmailUser يحتوي على حقل Email (string).
  ​أنشئ Struct اسمه SMSUser يحتوي على حقل PhoneNumber (string).

- ​تطبيق الـ Interface (Implementation):
  ​اجعل EmailUser يطبق الـ interface Notifier.

- ​التنفيذ: يطبع جملة "Sending Email to [Email]: [message]".

- ​اجعل SMSUser يطبق الـ interface Notifier.

- ​التنفيذ: يطبع جملة "Sending SMS to [PhoneNumber]: [message]".

- ​دالة موحدة (Polymorphic Function):
  ​أنشئ دالة عادية (Stand-alone function) خارج الـ structs اسمها Notify.

- ​المدخلات: تأخذ Notifier (كـ interface) وتأخذ message (string).

- ​الوظيفة: تستدعي دالة .Send() للكائن الذي تم تمريره لها.

- j# Main Function:
- ​أنشئ مستخدم Email: "yasser@gmail.com".
- ​أنشئ مستخدم SMS: "01000000000".

- ​استدعِ دالة Notify لإرسال رسالة "Welcome" لمستخدم الإيميل.
- ​استدعِ نفس الدالة Notify لإرسال رسالة "OTP: 1234" لمستخدم الـ SMS.

## 🚀 طريقة التشغيل

```bash
go run main.go
```

## ​🛑 المخرج المتوقع (Expected Output):

```Output
Sending Email to admin@company.com: Welcome
Sending SMS to 01000000000: OTP: 1234
```
