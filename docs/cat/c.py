
def binS(key):
    l,r=0,len(a)-1
    while(l<=r):
        mid=(l+r)//2
        midV=a[mid]
        if(midV==key):
            print('found',midV)
            break
        elif midV>key:
            r=mid-1
        elif midV<key:
            l=mid+1

# aの中のkey以上の最小のインデックス値を探索
def binSmin(key):
    ## 必ず lは key よりも小さい（条件を満たさない）
    ## 必ず rは key 以上
    l,r=-1,len(a)
    ## 探索範囲は lとrの差が1以下になるまで
    ## l<r だと rとlの差が1のときも条件が継続してしまう
    while(r-l>1):
        mid=(l+r)//2
        ## key以上
        if(key<=a[mid]):
            r=mid
        ## keyよりも小さい(条件を満たさない)
        else:
            l=mid

    if 0<r<len(a):
        print('min_key=',r,'minValue=',a[r])
        return r
    else:
        return -1

# aの中のkey以上の最大値を探索
def BinCondition(key,conditionFunc):
    ng,ok=-1,len(a)
    while(abs(ok-ng)>1):
        mid=(ok+ng)//2
        if conditionFunc(key):
            ok=mid
        else:
            ng=mid

    if 0<ok<len(a):
        print('min_key=',ok,'minValue=',a[ok])
        return ok
    else:
        return -1

    return
a=[1,2,3,4,4,4,4,4,4,4]

def binSmax(array,minIdx,key):
    ok,ng=minIdx,len(array)
    while abs(ok-ng)>1:
        mid=(ok+ng)//2
        if(array[mid]==key):
            ok=mid
        else:
            ng=mid
    if 0<ok<len(array):
        return ok
    else:
        return -1
    
print(binSmax(a,3,4),'ans=',len(a)-1)
# binSmin(9)
# print('年齢を当てます(10-15歳)まで')
# ok,ng=9,16
# while (abs(ok-ng)>1):
#     mid=(ok+ng)//2
#     key=input(f'{mid}才以上ですか？ Yes or No>>')
#     if key.lower()=='yes':
#         ok=mid
#     else:
#         ng=mid
# print('あなたは',ok,'才だ')