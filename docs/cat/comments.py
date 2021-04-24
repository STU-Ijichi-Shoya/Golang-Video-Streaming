import random


randomComments=[
    "ええやん",
    "なにこれ",
    "は?",
    "すっごーい",
    "wwwwwwww",
    "めっちゃかわいい",
    "すき",
    "良いアニメ",
    "謎すぎる",
    "test...",
    "かっこええ～",

    "ここの謎コメなんなんｗｗｗｗ",
    "自動生成コメントは草",
    "草",
    "コメントが的外れすぎる"
]
probabilityArray=[600,400,300,50,4,5,1]
random.shuffle(randomComments)
sumPro=sum(probabilityArray)
print('[')
for i in range(0,60*5,2):
    r=random.randint(0,sumPro+1)
    ss=0
    num=0
    for idx,s in enumerate(probabilityArray):
        if r>ss:
            num=idx
        ss+=s

    for j in range(num+1):
        print('{')
        print(f'\"comment\":\"{random.choice(randomComments)}\",\"time\":{i+j}')
        print('},')
print(']')
