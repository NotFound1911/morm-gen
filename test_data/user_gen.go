package testdata

import (
    "github.com/NotFound1911/morm"
    
    sqlx "database/sql"
    
)

const (
    UserName = "Name"
    UserAge = "Age"
    UserNickName = "NickName"
    UserPicture = "Picture"
)


func UserNameLT(val string) morm.Predicate {
    return morm.C("Name").LT(val)
}

func UserNameGT(val string) morm.Predicate {
    return morm.C("Name").GT(val)
}

func UserNameEQ(val string) morm.Predicate {
    return morm.C("Name").EQ(val)
}

func UserAgeLT(val *int) morm.Predicate {
    return morm.C("Age").LT(val)
}

func UserAgeGT(val *int) morm.Predicate {
    return morm.C("Age").GT(val)
}

func UserAgeEQ(val *int) morm.Predicate {
    return morm.C("Age").EQ(val)
}

func UserNickNameLT(val *sqlx.NullString) morm.Predicate {
    return morm.C("NickName").LT(val)
}

func UserNickNameGT(val *sqlx.NullString) morm.Predicate {
    return morm.C("NickName").GT(val)
}

func UserNickNameEQ(val *sqlx.NullString) morm.Predicate {
    return morm.C("NickName").EQ(val)
}

func UserPictureLT(val []byte) morm.Predicate {
    return morm.C("Picture").LT(val)
}

func UserPictureGT(val []byte) morm.Predicate {
    return morm.C("Picture").GT(val)
}

func UserPictureEQ(val []byte) morm.Predicate {
    return morm.C("Picture").EQ(val)
}


const (
    UserDetailAddress = "Address"
)


func UserDetailAddressLT(val string) morm.Predicate {
    return morm.C("Address").LT(val)
}

func UserDetailAddressGT(val string) morm.Predicate {
    return morm.C("Address").GT(val)
}

func UserDetailAddressEQ(val string) morm.Predicate {
    return morm.C("Address").EQ(val)
}
